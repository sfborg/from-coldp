package fcoldp

import (
	"github.com/gnames/coldp/ent/coldp"
	"github.com/sfborg/from-coldp/pkg/config"
)

type fcoldp struct {
	cfg config.Config
	c   coldp.Archive
}

func New(cfg config.Config) FromCoLDP {
	res := fcoldp{cfg: cfg}
	return &res
}

// GetCoLDP reads a CoLDP Archive from a file, preparing it for ingestion.
func (f *fcoldp) GetCoLDP(file string) (coldp.Archive, error) {
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
