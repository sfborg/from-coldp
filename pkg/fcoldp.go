package fcoldp

import (
	coldpConfig "github.com/gnames/coldp/config"
	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/coldp/io/arcio"
	"github.com/sfborg/from-coldp/internal/ent/sfgarc"
	"github.com/sfborg/from-coldp/pkg/config"
)

type fcoldp struct {
	cfg config.Config
	c   coldp.Archive
	s   sfgarc.Archive
}

func New(cfg config.Config, sfgarc sfgarc.Archive) FromCoLDP {
	res := fcoldp{
		cfg: cfg,
		s:   sfgarc,
	}
	return &res
}

// GetCoLDP reads a CoLDP Archive from a file, preparing it for ingestion.
func (fc *fcoldp) GetCoLDP(path string) (coldp.Archive, error) {
	cfg := coldpConfig.New()
	fc.c = arcio.New(cfg, path)
	err := fc.c.ResetCache()
	if err != nil {
		return nil, err
	}
	err = fc.c.Extract()
	if err != nil {
		return nil, err
	}
	err = fc.c.DirInfo()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ImportCoLDP converts a coldp.Archive to a Species File Group Archive
// database.
func (fc *fcoldp) ImportCoLDP(arc coldp.Archive) error {
	err := fc.importMeta()
	if err != nil {
		return err
	}
	return nil
}

// ExportSFGA writes a Species File Group Archive to a file.
func (fc *fcoldp) ExportSFGA(outputPath string) error {
	return nil
}
