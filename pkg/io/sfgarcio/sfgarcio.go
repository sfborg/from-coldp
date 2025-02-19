package sfgarcio

import (
	"database/sql"
	"errors"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/sfborg/from-coldp/pkg/ent/sfgarc"
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

func (s *sfgarcio) DB() *sql.DB {
	return s.db
}

func (s *sfgarcio) Exists() bool {
	if s.db == nil {
		return false
	}

	q := "SELECT id FROM version"

	var id string
	err := s.db.QueryRow(q).Scan(&id)
	if err != nil {
		slog.Error("Cannot get data from archive", "error", err)
		return false
	}
	if id == "" {
		slog.Error("Archive version is empty")
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
	if !s.Exists() {
		return errors.New("cannot find SFGA archive")
	}

	outPath = trimExtentions(outPath)

	// Perform the export
	err := s.sfdb.Export(outPath+".sql", false, s.cfg.WithZipOutput)
	if err != nil {
		return err
	}
	err = s.sfdb.Export(outPath+".sqlite", true, s.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}

func trimExtentions(outPath string) string {
	hasExt := false
	ext := filepath.Ext(outPath)
	var trimmed string
	if ext == ".zip" {
		hasExt = true
		outPath = strings.TrimSuffix(outPath, ext)
		trimmed += ext
		ext = filepath.Ext(outPath)
	}
	if ext == ".sql" || ext == ".sqlite" {
		hasExt = true
		outPath = strings.TrimSuffix(outPath, ext)
		trimmed = ext + trimmed
	}
	if hasExt {
		slog.Warn("Trimmed extentions from output File", "ext", trimmed)
	}
	return outPath
}
