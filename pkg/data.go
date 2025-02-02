package fcoldp

import (
	"github.com/gnames/coldp/ent/coldp"
	"github.com/sfborg/from-coldp/pkg/ent/sfgarc"
)

func (fc *fcoldp) importData(c coldp.Archive) error {
	var err error
	var hasRefs bool
	paths := c.DataPaths()

	if res, ok := paths[coldp.ReferenceDT]; ok {
		if err = importData(fc, res, c, insertReferences); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.ReferenceJsonDT]; ok && !hasRefs {
		if err = importData(fc, res, c, insertReferences); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.AuthorDT]; ok {
		if err = importData(fc, res, c, insertAuthors); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.NameDT]; ok {
		if err = importData(fc, res, c, insertNames); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.TaxonDT]; ok {
		if err = importData(fc, res, c, insertTaxa); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.SynonymDT]; ok {
		if err = importData(fc, res, c, insertSynonyms); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.NameUsageDT]; ok {
		if err = importData(fc, res, c, insertNameUsages); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.VernacularNameDT]; ok {
		if err = importData(fc, res, c, insertVernaculars); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.NameRelationDT]; ok {
		if err = importData(fc, res, c, insertNameRelations); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.TypeMaterialDT]; ok {
		if err = importData(fc, res, c, insertTypeMaterials); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.DistributionDT]; ok {
		if err = importData(fc, res, c, insertDistributions); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.MediaDT]; ok {
		if err = importData(fc, res, c, insertMedia); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.TreatmentDT]; ok {
		if err = importData(fc, res, c, insertTreatments); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.SpeciesEstimateDT]; ok {
		if err = importData(fc, res, c, insertSpeciesEstimates); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.TaxonPropertyDT]; ok {
		if err = importData(fc, res, c, insertTaxonProperties); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.SpeciesInteractionDT]; ok {
		if err = importData(
			fc, res, c, insertSpeciesInteractions,
		); err != nil {
			return err
		}
	}
	if res, ok := paths[coldp.TaxonConceptRelationDT]; ok {
		if err = importData(fc, res, c, insertTaxonConceptRels); err != nil {
			return err
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
