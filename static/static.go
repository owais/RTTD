// Package database contains the database artefacts of GOM as embedded resource
package static

import (
	"embed"
)

// // no go:generate parcello -r

//go:embed files/*
var Files embed.FS
