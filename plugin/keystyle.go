// This must be package main
package main

import (
	"fmt"

	"github.com/merschformann/keystyle/analyzer"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/tools/go/analysis"
)

type MapStyle struct {
	// Map is the map type to check.
	Map string `mapstructure:"map"`
	// Style is the style to adhere to.
	// It can be one of the following:
	// - camelCase
	// - PascalCase
	// - snake_case
	// - kebab-case
	Style string `mapstructure:"style"`
	// Ignore is a list of keys to ignore.
	Ignore []string `mapstructure:"ignore"`
}

type KeyStyleOptions struct {
	// Maps is the list of all map types to check.
	Maps []MapStyle `mapstructure:"maps"`
}

func New(conf any) ([]*analysis.Analyzer, error) {
	// TODO: This must be implemented

	fmt.Printf("My configuration (%[1]T): %#[1]v\n", conf)

	var config KeyStyleOptions
	err := mapstructure.Decode(conf, &config)
	if err != nil {
		panic(err)
	}

	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	return []*analysis.Analyzer{analyzer.KeyStyleAnalyzer}, nil
}
