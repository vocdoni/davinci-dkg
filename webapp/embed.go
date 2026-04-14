// Package webapp embeds the built DKG explorer single-page application so the
// dkg-node binary can serve it without any external assets.
package webapp

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var distFS embed.FS

// Assets returns the built webapp filesystem rooted at dist/.
func Assets() (fs.FS, error) {
	return fs.Sub(distFS, "dist")
}
