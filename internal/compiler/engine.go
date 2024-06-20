package compiler

import (
	"context"
	"fmt"

	"github.com/darkframemaster/sqlc/internal/analyzer"
	"github.com/darkframemaster/sqlc/internal/config"
	"github.com/darkframemaster/sqlc/internal/engine/dolphin"
	"github.com/darkframemaster/sqlc/internal/engine/postgresql"
	pganalyze "github.com/darkframemaster/sqlc/internal/engine/postgresql/analyzer"
	"github.com/darkframemaster/sqlc/internal/engine/sqlite"
	"github.com/darkframemaster/sqlc/internal/opts"
	"github.com/darkframemaster/sqlc/internal/quickdb"
	pb "github.com/darkframemaster/sqlc/internal/quickdb/v1"
	"github.com/darkframemaster/sqlc/internal/sql/catalog"
)

type Compiler struct {
	conf     config.SQL
	combo    config.CombinedSettings
	catalog  *catalog.Catalog
	parser   Parser
	result   *Result
	analyzer analyzer.Analyzer
	client   pb.QuickClient

	schema []string
}

func NewCompiler(conf config.SQL, combo config.CombinedSettings) (*Compiler, error) {
	c := &Compiler{conf: conf, combo: combo}

	if conf.Database != nil && conf.Database.Managed {
		client, err := quickdb.NewClientFromConfig(combo.Global.Cloud)
		if err != nil {
			return nil, fmt.Errorf("client error: %w", err)
		}
		c.client = client
	}

	switch conf.Engine {
	case config.EngineSQLite:
		c.parser = sqlite.NewParser()
		c.catalog = sqlite.NewCatalog()
	case config.EngineMySQL:
		c.parser = dolphin.NewParser()
		c.catalog = dolphin.NewCatalog()
	case config.EnginePostgreSQL:
		c.parser = postgresql.NewParser()
		c.catalog = postgresql.NewCatalog()
		if conf.Database != nil {
			if conf.Analyzer.Database == nil || *conf.Analyzer.Database {
				c.analyzer = analyzer.Cached(
					pganalyze.New(c.client, *conf.Database),
					combo.Global,
					*conf.Database,
				)
			}
		}
	default:
		return nil, fmt.Errorf("unknown engine: %s", conf.Engine)
	}
	return c, nil
}

func (c *Compiler) Catalog() *catalog.Catalog {
	return c.catalog
}

func (c *Compiler) ParseCatalog(schema []string) error {
	return c.parseCatalog(schema)
}

func (c *Compiler) ParseQueries(queries []string, o opts.Parser) error {
	r, err := c.parseQueries(o)
	if err != nil {
		return err
	}
	c.result = r
	return nil
}

func (c *Compiler) Result() *Result {
	return c.result
}

func (c *Compiler) Close(ctx context.Context) {
	if c.analyzer != nil {
		c.analyzer.Close(ctx)
	}
}
