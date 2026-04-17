package linters

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestCamelCase checks camelCase enforcement on string literals, const keys,
// and var keys.
func TestCamelCase(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "camelCase",
				"type-name": "LogData",
				"regex":     "",
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/keystyle")
}

// TestPascalCase checks PascalCase enforcement on string literals and const
// keys.
func TestPascalCase(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "PascalCase",
				"type-name": "EventData",
				"regex":     "",
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/pascalstyle")
}

// TestKebabCase checks kebab-case enforcement on string literals and const
// keys.
func TestKebabCase(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "kebab-case",
				"type-name": "HeaderData",
				"regex":     "",
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/kebabstyle")
}

// TestSnakeCase checks snake_case enforcement on string literals and const
// keys.
func TestSnakeCase(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "snake_case",
				"type-name": "MetricData",
				"regex":     "",
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/snakestyle")
}

// TestCustomStyle checks a custom regex style on string literals and const
// keys.
func TestCustomStyle(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "custom",
				"type-name": "CustomData",
				"regex":     `^key_[a-z]+$`,
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/customstyle")
}

// TestInvalidCustomStyleRegex ensures that an invalid regex causes a run-time
// error.
func TestInvalidCustomStyleRegex(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{"style": "custom", "type-name": "CustomData", "regex": `[invalid`},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)

	// The analyzer's Run function should return an error for a bad regex.
	_, runErr := analyzers[0].Run(nil)
	require.Error(t, runErr)
	require.Contains(t, runErr.Error(), "invalid regex for custom style")
}

// TestMissingCustomStyleRegex ensures that a custom style with an empty regex
// causes a run-time error.
func TestMissingCustomStyleRegex(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{"style": "custom", "type-name": "CustomData", "regex": ""},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)

	_, runErr := analyzers[0].Run(nil)
	require.Error(t, runErr)
	require.Contains(t, runErr.Error(), "custom style requires a regex pattern")
}

// TestUnknownStyle ensures that an unknown style name causes a run-time error.
func TestUnknownStyle(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{"style": "unknownStyle", "type-name": "LogData", "regex": ""},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	require.Len(t, analyzers, 1)

	_, runErr := analyzers[0].Run(nil)
	require.Error(t, runErr)
	require.Contains(t, runErr.Error(), "unknown style: unknownStyle")
}

func testdataDir(t *testing.T) string {
	t.Helper()

	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		require.Fail(t, "unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
