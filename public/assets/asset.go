package assets

import rice "github.com/GeertJohan/go.rice"

var AssetBox *rice.Box

// LoadAssets load asset box
func LoadAssets(path string) {
	AssetBox = rice.MustFindBox(path)
}
