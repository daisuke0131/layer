package layer_test

import (
	"fmt"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
	"layer"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	fmt.Println(testdata)
	analysistest.Run(t, testdata, layer.Analyzer,"a/...")
}
