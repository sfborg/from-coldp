package fcoldp

import (
	"github.com/gnames/coldp/ent/coldp"
	"github.com/sfborg/from-coldp/internal/ent/sfgarc"
)

func (fc *fcoldp) importData(c coldp.Archive) error {
	var err error
	paths := c.DataPaths()
	for k, v := range paths {
		switch k {
		case coldp.AuthorDT:
			if err = importDataGeneric(fc, v, c, insertAuthors); err != nil {
				return err
			}
		case coldp.DistributionDT:
			if err = importDataGeneric(fc, v, c, insertDistributions); err != nil {
				return err
			}
		case coldp.MediaDT:
			if err = importDataGeneric(fc, v, c, insertMedia); err != nil {
				return err
			}
		case coldp.NameDT:
			if err = importDataGeneric(fc, v, c, insertNames); err != nil {
				return err
			}
		case coldp.NameRelationDT:
			if err = importDataGeneric(fc, v, c, insertNameRelations); err != nil {
				return err
			}
		case coldp.NameUsageDT:
			if err = importDataGeneric(fc, v, c, insertNameUsages); err != nil {
				return err
			}
		case coldp.ReferenceDT:
			if err = importDataGeneric(fc, v, c, insertReferences); err != nil {
				return err
			}
		case coldp.SpeciesEstimateDT:
			if err = importDataGeneric(fc, v, c, insertSpeciesEstimates); err != nil {
				return err
			}
		case coldp.SpeciesInteractionDT:
			if err = importDataGeneric(
				fc, v, c, insertSpeciesInteractions,
			); err != nil {
				return err
			}
		case coldp.SynonymDT:
			if err = importDataGeneric(fc, v, c, insertSynonyms); err != nil {
				return err
			}
		case coldp.TaxonDT:
			if err = importDataGeneric(fc, v, c, insertTaxa); err != nil {
				return err
			}
		case coldp.TaxonConceptRelationDT:
			if err = importDataGeneric(fc, v, c, insertTaxonConceptRels); err != nil {
				return err
			}
		case coldp.TaxonPropertyDT:
			if err = importDataGeneric(fc, v, c, insertTaxonProperties); err != nil {
				return err
			}
		case coldp.TreatmentDT:
			if err = importDataGeneric(fc, v, c, insertTreatments); err != nil {
				return err
			}
		case coldp.TypeMaterialDT:
			if err = importDataGeneric(fc, v, c, insertTypeMaterials); err != nil {
				return err
			}
		case coldp.VernacularNameDT:
			if err = importDataGeneric(fc, v, c, insertVernaculars); err != nil {
				return err
			}
		}
	}
	return nil
}

func insertAuthors(s sfgarc.Archive, data []coldp.Author) error {
	return s.InsertAuthors(data)
}

func insertDistributions(s sfgarc.Archive, data []coldp.Distribution) error {
	return s.InsertDistributions(data)
}

func insertMedia(s sfgarc.Archive, data []coldp.Media) error {
	return s.InsertMedia(data)
}

func insertNames(s sfgarc.Archive, data []coldp.Name) error {
	return s.InsertNames(data)
}

func insertNameRelations(s sfgarc.Archive, data []coldp.NameRelation) error {
	return s.InsertNameRelations(data)
}

func insertNameUsages(s sfgarc.Archive, data []coldp.NameUsage) error {
	return s.InsertNameUsages(data)
}

func insertReferences(s sfgarc.Archive, data []coldp.Reference) error {
	return s.InsertReferences(data)
}

func insertSpeciesEstimates(
	s sfgarc.Archive,
	data []coldp.SpeciesEstimate,
) error {
	return s.InsertSpeciesEstimates(data)
}

func insertSpeciesInteractions(
	s sfgarc.Archive,
	data []coldp.SpeciesInteraction,
) error {
	return s.InsertSpeciesInteractions(data)
}

func insertSynonyms(s sfgarc.Archive, data []coldp.Synonym) error {
	return s.InsertSynonyms(data)
}

func insertTaxa(s sfgarc.Archive, data []coldp.Taxon) error {
	return s.InsertTaxa(data)
}

func insertTaxonConceptRels(
	s sfgarc.Archive,
	data []coldp.TaxonConceptRelation,
) error {
	return s.InsertTaxonConceptRelations(data)
}

func insertTaxonProperties(
	s sfgarc.Archive,
	data []coldp.TaxonProperty,
) error {
	return s.InsertTaxonProperties(data)
}

func insertTreatments(s sfgarc.Archive, data []coldp.Treatment) error {
	return s.InsertTreatments(data)
}

func insertTypeMaterials(s sfgarc.Archive, data []coldp.TypeMaterial) error {
	return s.InsertTypeMaterials(data)
}

func insertVernaculars(s sfgarc.Archive, data []coldp.Vernacular) error {
	return s.InsertVernaculars(data)
}
