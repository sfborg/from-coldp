package fcoldp

import (
	"github.com/gnames/coldp/ent/coldp"
)

func (fc *fcoldp) importData(c coldp.Archive) error {
	var err error
	paths := c.DataPaths()
	for k, v := range paths {
		switch k {
		case coldp.NameDT:
			err = fc.importName(v, c)
			if err != nil {
				return err
			}
		case coldp.TaxonDT:
			err = fc.importTaxon(v, c)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
