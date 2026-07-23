package skillissue

import "embed"

//go:embed all:dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills all:supporting-skills/dictate-plan all:evaluations/scenario-skill-refinement/*/skill all:evaluations/skill-calling/built-ins
var CanonicalSkills embed.FS
