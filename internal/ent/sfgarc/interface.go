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
	InsertSpeciesInteractions(data []coldp.SpeciesInteraction) error
	InsertSynonyms(data []coldp.Synonym) error
	InsertTaxa(data []coldp.Taxon) error
	InsertTaxonConceptRelations(data []coldp.TaxonConceptRelation) error
	InsertTaxonProperties(data []coldp.TaxonProperty) error
	InsertTreatments(data []coldp.Treatment) error
	InsertTypeMaterials(data []coldp.TypeMaterial) error
	InsertVernaculars(data []coldp.Vernacular) error

	Export(outPath string) error
}

type DataWriter interface {
	Write([]DataWriter) error
}
