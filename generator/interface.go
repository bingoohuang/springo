package generator

import (
	"github.com/bingoohuang/springo/generator/annotation"
	"github.com/bingoohuang/springo/model"
)

const (
	GenfilePrefix       = "gen_"
	GenfileExcludeRegex = GenfilePrefix + ".*"
)

type Generator interface {
	GetAnnotations() []annotation.AnnotationDescriptor
	Generate(inputDir string, parsedSources model.ParsedSources) error
}
