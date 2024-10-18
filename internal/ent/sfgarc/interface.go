package sfgarc

import "github.com/gnames/coldp/ent/coldp"

type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	InsertMeta(meta *coldp.Meta) error
	InsertNames(data []coldp.Name) error
	InsertTaxa(data []coldp.Taxon) error
	InsertSynonyms(data []coldp.Synonym) error
	InsertVernaculars(data []coldp.Vernacular) error
	InsertReferences(data []coldp.Reference) error

	Export(outPath string) error
}

type DataWriter interface {
	Write([]DataWriter) error
}
