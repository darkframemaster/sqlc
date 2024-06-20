// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/darkframemaster/sqlc/internal/sql/ast"
	"github.com/darkframemaster/sqlc/internal/sql/catalog"
)

var funcsPgrowlocks = []*catalog.Function{
	{
		Name: "pgrowlocks",
		Args: []*catalog.Argument{
			{
				Name: "relname",
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
}

func Pgrowlocks() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgrowlocks
	return s
}
