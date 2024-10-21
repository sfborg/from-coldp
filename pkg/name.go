package fcoldp

import (
	"context"

	"github.com/gnames/coldp/ent/coldp"
	"golang.org/x/sync/errgroup"
)

func (fc *fcoldp) importName(path string, c coldp.Archive) error {
	chIn := make(chan coldp.Name)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		err := fc.processNames(chIn)
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

func (fc *fcoldp) processNames(chIn <-chan coldp.Name) error {
	var err error
	names := make([]coldp.Name, 0, fc.cfg.BatchSize)
	var count int

	for n := range chIn {
		count++
		names = append(names, n)
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
