package dist

import "embed"

//go:embed dist/*
var DistFS embed.FS
