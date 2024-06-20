// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/darkframemaster/sqlc/internal/sql/ast"
	"github.com/darkframemaster/sqlc/internal/sql/catalog"
)

var funcsBtreeGist = []*catalog.Function{
	{
		Name: "cash_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "money"},
			},
			{
				Type: &ast.TypeName{Name: "money"},
			},
		},
		ReturnType: &ast.TypeName{Name: "money"},
	},
	{
		Name: "date_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "date"},
			},
			{
				Type: &ast.TypeName{Name: "date"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "float4_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "real"},
			},
			{
				Type: &ast.TypeName{Name: "real"},
			},
		},
		ReturnType: &ast.TypeName{Name: "real"},
	},
	{
		Name: "float8_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
			{
				Type: &ast.TypeName{Name: "double precision"},
			},
		},
		ReturnType: &ast.TypeName{Name: "double precision"},
	},
	{
		Name: "gbtreekey16_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey16"},
	},
	{
		Name: "gbtreekey16_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey16"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "gbtreekey2_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey2"},
	},
	{
		Name: "gbtreekey2_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey2"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "gbtreekey32_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey32"},
	},
	{
		Name: "gbtreekey32_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey32"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "gbtreekey4_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey4"},
	},
	{
		Name: "gbtreekey4_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey4"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "gbtreekey8_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey8"},
	},
	{
		Name: "gbtreekey8_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey8"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "gbtreekey_var_in",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "gbtreekey_var"},
	},
	{
		Name: "gbtreekey_var_out",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "gbtreekey_var"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "int2_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "smallint"},
			},
			{
				Type: &ast.TypeName{Name: "smallint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "smallint"},
	},
	{
		Name: "int4_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "integer"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "int8_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bigint"},
	},
	{
		Name: "interval_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "interval"},
			},
			{
				Type: &ast.TypeName{Name: "interval"},
			},
		},
		ReturnType: &ast.TypeName{Name: "interval"},
	},
	{
		Name: "oid_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "oid"},
			},
			{
				Type: &ast.TypeName{Name: "oid"},
			},
		},
		ReturnType: &ast.TypeName{Name: "oid"},
	},
	{
		Name: "time_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "time without time zone"},
			},
			{
				Type: &ast.TypeName{Name: "time without time zone"},
			},
		},
		ReturnType: &ast.TypeName{Name: "interval"},
	},
	{
		Name: "ts_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "timestamp without time zone"},
			},
			{
				Type: &ast.TypeName{Name: "timestamp without time zone"},
			},
		},
		ReturnType: &ast.TypeName{Name: "interval"},
	},
	{
		Name: "tstz_dist",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "timestamp with time zone"},
			},
			{
				Type: &ast.TypeName{Name: "timestamp with time zone"},
			},
		},
		ReturnType: &ast.TypeName{Name: "interval"},
	},
}

func BtreeGist() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsBtreeGist
	return s
}
