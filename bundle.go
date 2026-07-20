package skillissue

import "embed"

//go:embed all:skills all:supporting-skills all:evaluations/skill-calling/built-ins
var CanonicalSkills embed.FS
