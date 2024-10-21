package fcoldp

import (
	"context"

	"github.com/gnames/coldp/ent/coldp"
	"golang.org/x/sync/errgroup"
)

func (fc *fcoldp) importTaxon(path string, c coldp.Archive) error {
	chIn := make(chan coldp.Taxon)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		err := fc.processTaxa(chIn)
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

func (fc *fcoldp) processTaxa(chIn <-chan coldp.Taxon) error {
	var err error
	taxa := make([]coldp.Taxon, 0, fc.cfg.BatchSize)
	var count int

	for n := range chIn {
		count++
		taxa = append(taxa, n)
		if count == fc.cfg.BatchSize {
			err = fc.s.InsertTaxa(taxa)
			count = 0
			taxa = taxa[:0]
			if err != nil {
				return err
			}
		}
	}

	err = fc.s.InsertTaxa(taxa[:count])
	if err != nil {
		return err
	}

	return nil
}
