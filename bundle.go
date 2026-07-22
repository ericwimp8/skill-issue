package skillissue

import "embed"

//go:embed all:plugins/skill-issue/skills all:supporting-skills/dictate-plan all:evaluations/scenario-skill-refinement/*/skill all:evaluations/skill-calling/built-ins
var CanonicalSkills embed.FS
