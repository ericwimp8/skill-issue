package skillissue

import "embed"

//go:embed all:skills all:supporting-skills all:evaluation-skills
var CanonicalSkills embed.FS
