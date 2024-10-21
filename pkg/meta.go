package fcoldp

import "github.com/gnames/coldp/ent/coldp"

func (fc *fcoldp) importMeta(c coldp.Archive) error {
	meta, err := c.Meta()
	if err != nil {
		return err
	}

	err = fc.s.InsertMeta(meta)
	if err != nil {
		return err
	}

	return nil
}
