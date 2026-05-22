package lotp

import (
	"slices"
	"testing"
)

func TestLoad(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(c.Tools) == 0 {
		t.Fatal("catalog has no tools")
	}
	if c.Source == "" || c.SourceRef == "" {
		t.Errorf("catalog provenance missing: source=%q ref=%q", c.Source, c.SourceRef)
	}
	validCategory := map[string]bool{"config-file": true, "input-file": true, "env-var": true}
	for _, tool := range c.Tools {
		if tool.Name == "" {
			t.Errorf("tool with empty name: %+v", tool)
		}
		if len(tool.Categories) == 0 {
			t.Errorf("tool %q has no categories", tool.Name)
		}
		for _, cat := range tool.Categories {
			if !validCategory[cat] {
				t.Errorf("tool %q has unexpected category %q", tool.Name, cat)
			}
		}
	}
}

// TestMatcherToolsIn is the concrete match / no-match corpus for the LOTP
// detector: shell commands that must surface a tool and commands that must
// stay silent.
func TestMatcherToolsIn(t *testing.T) {
	c, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	m := c.Matcher()

	tests := []struct {
		name   string
		script string
		want   []string // expected tool names, sorted
	}{
		// --- must match: real RCE-by-design invocations ---
		{"npx eslint", "npx eslint .", []string{"eslint", "npx"}},
		{"npm lifecycle", "npm test", []string{"npm"}},
		{"make runs Makefile", "make build", []string{"make"}},
		{"gradle build script", "gradle build", []string{"gradle"}},
		{"pip reads requirements", "pip install -r requirements.txt", []string{"pip"}},
		{"pytest loads conftest", "pytest tests/", []string{"pytest"}},
		{"maven override mvn", "mvn -B verify", []string{"mvn"}},
		{"hyphenated tool name", "golangci-lint run ./...", []string{"golangci-lint"}},
		{"multiline script", "npm ci\nnpm run build", []string{"npm"}},
		{"chained commands", "make deps && pytest", []string{"make", "pytest"}},

		// --- must NOT match: inert commands ---
		{"echo", "echo hello world", nil},
		{"ls", "ls -la", nil},
		{"git clone", "git clone https://github.com/x/y", nil},
		{"cmake is not make", "cmake --build build", nil},
		{"composed name is not eslint", "cat eslint-config-airbnb.json", nil},
		{"curl", "curl -sSL https://example.com", nil},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got []string
			for _, tool := range m.ToolsIn(tc.script) {
				got = append(got, tool.Name)
			}
			if !slices.Equal(got, tc.want) {
				t.Errorf("ToolsIn(%q) = %v, want %v", tc.script, got, tc.want)
			}
		})
	}
}
