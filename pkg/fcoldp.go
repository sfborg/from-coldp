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
	c := arcio.New(cfg, path)
	// Resets cache for coldp working dir
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
	return c, nil
}

// ImportCoLDP converts a coldp.Archive to a Species File Group Archive
// database.
func (fc *fcoldp) ImportCoLDP(c coldp.Archive) error {
	err := fc.importMeta(c)
	if err != nil {
		return err
	}

	err = fc.importData(c)
	if err != nil {
		return err
	}
	return nil
}

// ExportSFGA writes a Species File Group Archive to a file.
func (fc *fcoldp) ExportSFGA(outputPath string) error {
	return nil
}
