package layer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"strconv"
	"strings"
)

const doc = "layer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "layer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files{
		dir := pass.Fset.File(f.Pos()).Name()
		if strings.Contains(dir,"/usecase/service/"){
			// infrastructureの層への依存はしない
			for _,imp := range  f.Imports{
				v, err := strconv.Unquote(imp.Path.Value)
				if err != nil{
					return nil, err
				}
				if strings.Contains(v,"/infrastructure/") {
					pass.Reportf(imp.Pos(), "%s is invalid import.", v)
				}
			}
		}
	}

	return nil, nil
}
