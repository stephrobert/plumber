# lotp-untrusted-code — flag a developer tool with RCE-by-design behaviour
# (Living Off The Pipeline) that runs on untrusted code.
#
# The catalogued CLIs (eslint, npm, make, gradle, pip, pytest, ...) execute
# attacker-controlled repository files when they run: eslint loads
# eslint.config.js, npm runs package.json lifecycle scripts, make runs the
# Makefile, and so on. When such a tool processes a fork's code under a
# privileged trigger, it yields arbitrary code execution with the base
# repository's secrets in scope.
#
# This rule reports a job only when all three hold:
#   1. it is reachable from a trigger that exposes base-repo secrets to an
#      unprivileged caller (pull_request_target, workflow_run);
#   2. it checks out the PR head (attacker-controlled code);
#   3. a later run step invokes a LOTP catalog tool.
#
# Inter-step taint: the tool must appear AFTER the untrusted checkout in the
# same job. A run step that precedes the head checkout processes the trusted
# base-repo code and is not flagged — `input.pipeline.jobs[].steps` carries
# the source order the collector preserved for exactly this reason.
package lotp_untrusted_code

import rego.v1

# Triggers that run with the base repository's secrets while being
# influenceable by an unprivileged caller.
dangerous_triggers := {"pull_request_target", "workflow_run"}

# Refs that resolve to attacker-controlled content under those triggers.
fork_ref_patterns := [
	`github\.event\.pull_request\.head\.sha`,
	`github\.event\.pull_request\.head\.ref`,
	`github\.event\.workflow_run\.head_sha`,
	`github\.event\.workflow_run\.head_branch`,
	`github\.head_ref`,
]

# Patterns that, in an `if:` guard, restrict the job to same-repository
# (non-fork) pull requests — the fork's code never runs, so the LOTP tool
# only ever processes trusted code. The canonical "Svelte pattern".
fork_guard_patterns := [
	`head\.repo\.full_name\s*==\s*github\.repository`,
	`github\.repository\s*==\s*[a-z._]*head\.repo\.full_name`,
	`head\.repo\.fork\s*==\s*false`,
	`!\s*[a-z._]*head\.repo\.fork\b`,
]

deny contains finding if {
	some i
	job := input.pipeline.jobs[i]
	_under_dangerous_trigger(job)
	not _has_fork_guard(job)

	# the untrusted checkout, at step index `c`
	some c
	_is_untrusted_checkout(job.steps[c])

	# a LOTP tool invoked by a run step AFTER that checkout
	some r
	r > c
	step := job.steps[r]
	step.kind == "run"
	count(step.lotpTools) > 0

	finding := {
		"code":      "ISSUE-414",
		"severity":  "critical",
		"message":   sprintf("job %q runs the Living Off The Pipeline tool(s) %s on untrusted code — a privileged trigger checks out the PR head, then this step lets attacker-controlled repository files execute with base-repo secrets in scope", [job.name, concat(", ", step.lotpTools)]),
		"job":       job.name,
		"lotpTools": step.lotpTools,
	}
}

_under_dangerous_trigger(job) if {
	some t in job.triggers
	dangerous_triggers[t]
}

_is_untrusted_checkout(step) if {
	step.kind == "uses"
	startswith(step.uses, "actions/checkout@")
	ref := step.with.ref
	is_string(ref)
	some p in fork_ref_patterns
	regex.match(p, ref)
}

# _has_fork_guard is true when any `if:` condition on the job or its steps
# limits execution to same-repository pull requests.
_has_fork_guard(job) if {
	some cond in job.conditions
	some p in fork_guard_patterns
	regex.match(p, cond)
}
