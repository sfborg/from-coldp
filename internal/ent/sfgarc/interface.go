package sfgarc

import "github.com/gnames/coldp/ent/coldp"

type Archive interface {
	Exists() bool
	Connect() error
	Close() error
	InsertMeta(meta *coldp.Meta) error

	InsertAuthors(data []coldp.Author) error
	InsertDistributions(data []coldp.Distribution) error
	InsertMedia(data []coldp.Media) error
	InsertNames(data []coldp.Name) error
	InsertNameRelations(data []coldp.NameRelation) error
	InsertNameUsages(data []coldp.NameUsage) error
	InsertReferences(data []coldp.Reference) error
	InsertSpeciesEstimates(data []coldp.SpeciesEstimate) error
	InsertSynonyms(data []coldp.Synonym) error
	InsertTaxa(data []coldp.Taxon) error
	InsertVernaculars(data []coldp.Vernacular) error

	Export(outPath string) error
}

type DataWriter interface {
	Write([]DataWriter) error
}
