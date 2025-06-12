package linters

import (
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("keystyle", New)
}

type KeyStyleSettings struct {
	Checks []KeyStyleDefinition `json:"checks"` // List of key style definitions to check.
}

type KeyStyleDefinition struct {
	Style    KeyStyle `json:"style"`     // The style of keys in maps, e.g., "camelCase", "PascalCase", "kebab-case", "snake_case", "custom".
	TypeName string   `json:"type-name"` // The name of the type of map to check, e.g., "LogData"
	Regex    string   `json:"regex"`     // Optional custom regex to match keys (if Style is "custom").
}

type PluginKeyStyle struct {
	settings KeyStyleSettings
}

func New(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	s, err := register.DecodeSettings[KeyStyleSettings](settings)
	if err != nil {
		return nil, err
	}

	return &PluginKeyStyle{settings: s}, nil
}

func (f *PluginKeyStyle) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "keystyle",
			Doc:  "Checks the style of keys in maps.",
			Run:  f.run,
		},
	}, nil
}

func (f *PluginKeyStyle) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.
	return register.LoadModeSyntax
}

// KeyStyle defines the style of keys in maps.
type KeyStyle string

const (
	CamelCase  KeyStyle = "camelCase"
	PascalCase KeyStyle = "PascalCase"
	KebabCase  KeyStyle = "kebab-case"
	SnakeCase  KeyStyle = "snake_case"
	Custom     KeyStyle = "custom" // Custom style with a regex pattern.
)

var camelCaseRegex = regexp.MustCompile("^[a-z]+[0-9]*([A-Z][a-z]*[0-9]*)*$")
var pascalCaseRegex = regexp.MustCompile("^([A-Z][a-z0-9]+)+$")
var kebabCaseRegex = regexp.MustCompile("^[a-z]+(-[a-z0-9]+)*$")
var snakeCaseRegex = regexp.MustCompile("^[a-z]+(_[a-z0-9]+)*$")

func checkStyle(s string, style *regexp.Regexp) bool {
	return style.MatchString(s)
}

func (f *PluginKeyStyle) run(pass *analysis.Pass) (interface{}, error) {

	for _, styleCheck := range f.settings.Checks {
		var styleRegex *regexp.Regexp
		switch styleCheck.Style {
		case CamelCase:
			styleRegex = camelCaseRegex
		case PascalCase:
			styleRegex = pascalCaseRegex
		case KebabCase:
			styleRegex = kebabCaseRegex
		case SnakeCase:
			styleRegex = snakeCaseRegex
		case Custom:
			if styleCheck.Regex == "" {
				return nil, fmt.Errorf("custom style requires a regex pattern")
			}
			var err error
			styleRegex, err = regexp.Compile(styleCheck.Regex)
			if err != nil {
				return nil, fmt.Errorf("invalid regex for custom style: %w", err)
			}
		default:
			return nil, fmt.Errorf("unknown style: %s", styleCheck.Style)
		}

		for _, file := range pass.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				// Check for composite literals
				if cl, ok := n.(*ast.CompositeLit); ok {
					// Check if the composite literal is of type CustomType or *CustomType or logger.CustomType
					var isCustomType bool
					switch t := cl.Type.(type) {
					case *ast.Ident:
						isCustomType = styleCheck.TypeName == t.Name
					case *ast.StarExpr:
						if ident, ok := t.X.(*ast.Ident); ok {
							isCustomType = styleCheck.TypeName == ident.Name
						} else if sel, ok := t.X.(*ast.SelectorExpr); ok {
							isCustomType = styleCheck.TypeName == sel.Sel.Name
						}
					case *ast.SelectorExpr:
						isCustomType = styleCheck.TypeName == t.Sel.Name
					default:
						isCustomType = false
					}

					if isCustomType {
						// Check each key in the composite literal
						for _, elt := range cl.Elts {
							if kv, ok := elt.(*ast.KeyValueExpr); ok {
								if keyIdent, ok := kv.Key.(*ast.Ident); ok {
									if !checkStyle(keyIdent.Name, styleRegex) {
										pass.Report(analysis.Diagnostic{
											Pos:            keyIdent.Pos(),
											End:            0,
											Category:       "keystyle",
											Message:        fmt.Sprintf("Key '%s' style should be camelCase", keyIdent.Name),
											SuggestedFixes: nil,
										})
										return false // Stop inspecting this node
									}
								} else if keyBasicLit, ok := kv.Key.(*ast.BasicLit); ok && keyBasicLit.Kind == token.STRING {
									key := strings.Trim(keyBasicLit.Value, `"`)
									if !checkStyle(key, styleRegex) {
										pass.Report(analysis.Diagnostic{
											Pos:            keyBasicLit.Pos(),
											End:            0,
											Category:       "keystyle",
											Message:        fmt.Sprintf("Key '%s' style should be camelCase", key),
											SuggestedFixes: nil,
										})
										return false // Stop inspecting this node
									}
								}
							}
						}
					}
				}
				return true
			})
		}
	}

	return nil, nil
}
