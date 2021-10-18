package parser

import "github.com/bingoohuang/springo/model"

type Parser interface {
	ParseSourceDir(dirName string, includeRegex string, excludeRegex string) (model.ParsedSources, error)
}
