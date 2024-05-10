package createdb

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/singleflight"

	"github.com/sqlc-dev/sqlc/internal/pgx/poolcache"
)

type Server struct {
	uri    string
	flight singleflight.Group
	cache  *poolcache.Cache
}

func New(uri string) *Server {
	return &Server{
		uri:   uri,
		cache: poolcache.New(),
	}
}

func (s *Server) Create(ctx context.Context, hash string, migrations []string) (string, *pgxpool.Pool, error) {
	name := fmt.Sprintf("sqlc_%s", hash)

	uri, err := url.Parse(s.uri)
	if err != nil {
		return "", nil, err
	}
	uri.Path = name

	key := uri.String()
	_, err, _ = s.flight.Do(key, func() (interface{}, error) {
		// TODO: Use a parameterized query
		row := s.pool.QueryRow(ctx,
			fmt.Sprintf(`SELECT datname FROM pg_database WHERE datname = '%s'`, name))

		var datname string
		if err := row.Scan(&datname); err == nil {
			slog.Info("database exists", "name", name)
			return nil, nil
		}

		slog.Info("creating database", "name", name)
		if _, err := s.pool.Exec(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, name)); err != nil {
			return nil, err
		}

		conn, err := pgx.Connect(ctx, uri.String())
		if err != nil {
			return nil, fmt.Errorf("connect %s: %s", name, err)
		}
		defer conn.Close(ctx)

		for _, q := range migrations {
			if len(strings.TrimSpace(q)) == 0 {
				continue
			}
			if _, err := conn.Exec(ctx, q); err != nil {
				return nil, fmt.Errorf("%s: %s", q, err)
			}
		}
		return nil, nil
	})

	if err != nil {
		return "", nil, err
	}

	db, err := s.cache.Open(ctx, key)
	return key, db, err
}
