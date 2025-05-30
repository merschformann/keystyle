package linters

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTodoAnalyzer(t *testing.T) {
	linter, err := New(map[string]any{
		"checks": []map[string]any{
			{
				"style":     "camelCase", // Example style, can be changed to "PascalCase", "kebab-case", etc.
				"type-name": "LogData",   // Example type name to check.
				"regex":     "",          // Optional regex for custom key matching.
			},
		},
	})
	require.NoError(t, err)
	analyzers, err := linter.BuildAnalyzers()
	require.NoError(t, err)
	if len(analyzers) != 1 {
		require.Fail(t, "expected exactly one analyzer, got %d", len(analyzers))
	}
	analysistest.Run(t, testdataDir(t), analyzers[0], "testlintdata/keystyle")
}

func testdataDir(t *testing.T) string {
	t.Helper()

	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		require.Fail(t, "unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
