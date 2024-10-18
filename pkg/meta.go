package fcoldp

func (fc *fcoldp) importMeta() error {
	meta, err := fc.c.Meta()
	if err != nil {
		return err
	}

	err = fc.s.InsertMeta(meta)
	if err != nil {
		return err
	}

	return nil
}
