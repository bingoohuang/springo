package ast

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bingoohuang/springo/generator"
	"github.com/bingoohuang/springo/generator/annotation"
	"github.com/bingoohuang/springo/generator/event/eventAnnotation"
	"github.com/bingoohuang/springo/generator/generationUtil"
	"github.com/bingoohuang/springo/model"
)

type Generator struct {
}

func NewGenerator() generator.Generator {
	return &Generator{}
}

func (eg *Generator) GetAnnotations() []annotation.AnnotationDescriptor {
	return eventAnnotation.Get()
}

func (eg *Generator) Generate(inputDir string, parsedSources model.ParsedSources) error {
	marshalled, err := json.MarshalIndent(parsedSources, "", "\t")
	if err != nil {
		panic(err)
	}
	targetFilename := generationUtil.Prefixed(inputDir + "/" + "ast.json")
	err = ioutil.WriteFile(targetFilename, marshalled, 0o644)
	if err != nil {
		panic(err)
	}

	return nil
}
