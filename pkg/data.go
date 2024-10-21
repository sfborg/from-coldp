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
		case coldp.NameDT:
			if err = importDataGeneric(fc, v, c, insertNames); err != nil {
				return err
			}
		case coldp.TaxonDT:
			if err = importDataGeneric(fc, v, c, insertTaxa); err != nil {
				return err
			}
		case coldp.SynonymDT:
			if err = importDataGeneric(fc, v, c, insertSynonym); err != nil {
				return err
			}
		}
	}
	return nil
}

func insertNames(s sfgarc.Archive, data []coldp.Name) error {
	return s.InsertNames(data)
}

func insertTaxa(s sfgarc.Archive, data []coldp.Taxon) error {
	return s.InsertTaxa(data)
}

func insertSynonym(s sfgarc.Archive, data []coldp.Synonym) error {
	return s.InsertSynonyms(data)
}
