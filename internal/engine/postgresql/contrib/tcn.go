// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/darkframemaster/sqlc/internal/sql/ast"
	"github.com/darkframemaster/sqlc/internal/sql/catalog"
)

var funcsTcn = []*catalog.Function{
	{
		Name:       "triggered_change_notification",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "trigger"},
	},
}

func Tcn() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsTcn
	return s
}
