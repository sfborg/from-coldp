package fcoldp

import (
	"context"
	"database/sql"

	"github.com/gnames/coldp/ent/coldp"
	"golang.org/x/sync/errgroup"
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

func dwNew(t coldp.DattaLoader, f func(*sql.DB, []T)) error {
	return dataWriter{data: t, f: f}
}

type dataWriter[T coldp.DataLoader] struct {
	data T
	f    func(*sql.DB, []T) error
}

func (dw dataWriter[T]) importData(fc *fcoldp, path string, c coldp.Archive) error {
	chIn := make(chan T)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		err := dw.process(fc, chIn)
		if err != nil {
			for range chIn {
			}
		}
		return err
	})

	err := coldp.Read(c.Config(), path, chIn)
	if err != nil {
		return err
	}
	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}

func (dw dataWriter[T]) process(fc *fcoldp, chIn chan T) error {
	return nil
}

func (fc *fcoldp) processNames(chIn <-chan coldp.Name) error {
	var err error
	names := make([]*coldp.Name, 0, fc.cfg.BatchSize)
	var count int

	for n := range chIn {
		count++
		names = append(names, &n)
		if count == fc.cfg.BatchSize {
			err = fc.s.InsertNames(names)
			count = 0
			names = names[:0]
			if err != nil {
				return err
			}
		}
	}

	err = fc.s.InsertNames(names[:count])
	if err != nil {
		return err
	}

	return nil
}
