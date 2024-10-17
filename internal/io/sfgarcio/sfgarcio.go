package sfgarcio

import (
	"database/sql"
	"log/slog"

	"github.com/sfborg/from-coldp/internal/ent/sfgarc"
	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type sfgarcio struct {
	cfg  config.Config
	sch  sfga.Schema
	sfdb sfga.DB
	db   *sql.DB
}

// New creates an instance of SFGArchive store
func New(cfg config.Config, sch sfga.Schema, sfdb sfga.DB) sfgarc.Archive {
	return &sfgarcio{cfg: cfg, sch: sch, sfdb: sfdb}
}

func (s *sfgarcio) Exists() bool {
	if s.db == nil {
		return false
	}

	q := "SELECT dwc_taxon_id FROM core LIMIT 5"

	var id string
	err := s.db.QueryRow(q).Scan(&id)
	if err != nil {
		slog.Error("Cannot get data from core", "error", err)
		return false
	}
	if id == "" {
		slog.Error("No dwc_taxon_id in core")
		return false
	}

	return true
}

func (s *sfgarcio) Close() error {
	if s.db == nil {
		return nil
	}
	return s.sfdb.Close()
}

func (s *sfgarcio) Export(outPath string) error {
	return nil
}
