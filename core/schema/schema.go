package schema

import (
	"github.com/liuzz1983/scalesearch/core/analysis"
	"github.com/liuzz1983/scalesearch/core/errors"
)

type SchemaField struct {
	Analyzer analysis.Analyzer
}

type Schema struct {
	fields map[string]*SchemaField
}

func (schema *Schema) Analyzer(fieldName string) (analysis.Analyzer, error) {
	fieldSchema, ok := schema.fields[fieldName]
	if !ok {
		return nil, errors.New("can not find field schema")
	}
	return fieldSchema.Analyzer, nil
}
