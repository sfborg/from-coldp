package sfgarc

import "github.com/gnames/coldp/ent/coldp"

type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	InsertMeta(meta *coldp.Meta) error
	InsertNames(names []coldp.Name) error
	InsertTaxa(names []coldp.Taxon) error
	Export(outPath string) error
}

type DataWriter interface {
	Write([]DataWriter) error
}
