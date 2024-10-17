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
func (f *fcoldp) GetCoLDP(path string) (coldp.Archive, error) {
	cfg := coldpConfig.New()
	c := arcio.New(cfg, path)
	err := c.ResetCache()
	if err != nil {
		return nil, err
	}
	err = c.Extract()
	if err != nil {
		return nil, err
	}
	err = c.DirInfo()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ImportCoLDP converts a coldp.Archive to a Species File Group Archive
// database.
func (f *fcoldp) ImportCoLDP(arc coldp.Archive) error {
	return nil
}

// ExportSFGA writes a Species File Group Archive to a file.
func (f *fcoldp) ExportSFGA(outputPath string) error {
	return nil
}
