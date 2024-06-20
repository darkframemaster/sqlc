// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/darkframemaster/sqlc/internal/sql/ast"
	"github.com/darkframemaster/sqlc/internal/sql/catalog"
)

var funcsPgVisibility = []*catalog.Function{
	{
		Name: "pg_check_frozen",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "tid"},
	},
	{
		Name: "pg_check_visible",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "tid"},
	},
	{
		Name: "pg_truncate_visibility_map",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
	{
		Name: "pg_visibility",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_visibility",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_visibility_map",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_visibility_map",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
			{
				Name: "blkno",
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_visibility_map_summary",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regclass"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
}

func PgVisibility() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgVisibility
	return s
}
