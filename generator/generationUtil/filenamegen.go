package generationUtil

import (
	"path"

	"github.com/bingoohuang/springo/generator"
)

func Prefixed(filenamePath string) string {
	dir, filename := path.Split(filenamePath)
	return dir + generator.GenfilePrefix + filename
}
