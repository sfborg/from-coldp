package fcoldp

import (
	"github.com/gnames/coldp/ent/coldp"
)

// FromCoLDP provies methods to convert CoLDP Archive to Species File Group
// Archive.
type FromCoLDP interface {
	// GetCoLDP reads a CoLDP Archive from a file, preparing it for ingestion.
	GetCoLDP(file string) (coldp.Archive, error)

	// ImportCoLDP converts a coldp.Archive to a Species File Group Archive
	// database.
	ImportCoLDP(arc coldp.Archive) error

	// ExportSFGA writes a Species File Group Archive to a file.
	ExportSFGA(outputPath string) error
}
