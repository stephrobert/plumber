// Package lotp exposes the Living Off The Pipeline (LOTP) catalog: developer
// CLI tools with RCE-by-design behaviour that execute attacker-supplied code
// when they process a repository file or a poisoned environment variable.
//
// The catalog is sourced from the open-source LOTP project
// (https://github.com/boostsecurityio/lotp, Apache-2.0) and embedded at build
// time. Refresh it with `go generate ./internal/lotp`.
//
// The catalog backs the LOTP control, which reports a tool only when it runs
// on untrusted code — see the lotp_untrusted_code Rego policy.
package lotp

//go:generate go run ./gen -out catalog.json

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"sync"
)

//go:embed catalog.json
var catalogJSON []byte

// Tool is one RCE-by-design CLI entry of the catalog.
type Tool struct {
	// Name is the executable name as it appears in a shell command.
	Name string `json:"name"`
	// Categories is the LOTP abuse tags carried by the tool, a subset of
	// {config-file, input-file, env-var}.
	Categories []string `json:"categories"`
	// Files lists the repository files (or globs) the tool loads and may
	// execute. Informational — used to enrich the finding message.
	Files []string `json:"files,omitempty"`
	// Reference is a documentation URL for the abuse vector.
	Reference string `json:"reference,omitempty"`
}

// Catalog is the embedded LOTP dataset.
type Catalog struct {
	Source    string `json:"source"`
	SourceRef string `json:"sourceRef"`
	License   string `json:"license"`
	Tools     []Tool `json:"tools"`
}

// Load decodes the embedded LOTP catalog.
func Load() (*Catalog, error) {
	var c Catalog
	if err := json.Unmarshal(catalogJSON, &c); err != nil {
		return nil, fmt.Errorf("decode lotp catalog: %w", err)
	}
	if len(c.Tools) == 0 {
		return nil, fmt.Errorf("lotp catalog is empty")
	}
	return &c, nil
}

// commandToken matches a shell-command word. The hyphen is kept inside the
// class so multi-segment names (golangci-lint, pre-commit) tokenise whole and
// flag names (--config) never collapse onto a bare tool name.
var commandToken = regexp.MustCompile(`[A-Za-z0-9_-]+`)

// Matcher detects LOTP tool invocations inside shell scripts.
type Matcher struct {
	byName map[string]Tool
}

// Matcher builds a Matcher indexed over the catalog's tools.
func (c *Catalog) Matcher() *Matcher {
	byName := make(map[string]Tool, len(c.Tools))
	for _, t := range c.Tools {
		byName[t.Name] = t
	}
	return &Matcher{byName: byName}
}

// ToolsIn returns every LOTP tool invoked in the given shell script,
// deduplicated and sorted by name. Matching is token-based: a tool is reported
// when its executable name appears as a whole word. It does not attempt full
// shell parsing — over-matching is acceptable because the LOTP control only
// fires once the surrounding step is reachable from untrusted code.
func (m *Matcher) ToolsIn(script string) []Tool {
	seen := make(map[string]Tool)
	for _, token := range commandToken.FindAllString(script, -1) {
		if tool, ok := m.byName[token]; ok {
			seen[tool.Name] = tool
		}
	}
	if len(seen) == 0 {
		return nil
	}
	out := make([]Tool, 0, len(seen))
	for _, t := range seen {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// ToolNamesIn is ToolsIn projected onto the tool names — the form the IR
// stores on a step.
func (m *Matcher) ToolNamesIn(script string) []string {
	tools := m.ToolsIn(script)
	if len(tools) == 0 {
		return nil
	}
	names := make([]string, len(tools))
	for i, t := range tools {
		names[i] = t.Name
	}
	return names
}

var (
	defaultOnce    sync.Once
	defaultMatcher *Matcher
	defaultErr     error
)

// DefaultMatcher returns a process-wide Matcher built once over the embedded
// catalog. It is safe for concurrent use.
func DefaultMatcher() (*Matcher, error) {
	defaultOnce.Do(func() {
		c, err := Load()
		if err != nil {
			defaultErr = err
			return
		}
		defaultMatcher = c.Matcher()
	})
	return defaultMatcher, defaultErr
}
