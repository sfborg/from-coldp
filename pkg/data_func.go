package fcoldp

import (
	"context"
	"path/filepath"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/sfborg/from-coldp/pkg/ent/sfgarc"
	"golang.org/x/sync/errgroup"
)

func importData[T coldp.DataLoader](
	fc *fcoldp,
	path string,
	c coldp.Archive,
	insertFunc func(sfgarc.Archive, []T) error) error {
	chIn := make(chan T)
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		return insert(fc.s, fc.cfg.BatchSize, chIn, insertFunc)
	})

	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		err = coldp.ReadJSON(path, chIn)
		if err != nil {
			return err
		}
	case ".jsonl":
		err = coldp.ReadJSONL(path, chIn)
		if err != nil {
			return err
		}
	default:
		err = coldp.Read(c.Config(), path, chIn)
		if err != nil {
			return err
		}
	}
	close(chIn)
	if err = g.Wait(); err != nil {
		return err
	}

	return nil
}

func insert[T coldp.DataLoader](
	s sfgarc.Archive,
	batchSize int,
	ch <-chan T,
	insertFunc func(sfgarc.Archive, []T) error,
) error {
	var err error
	names := make([]T, 0, batchSize)
	var count int

	for n := range ch {
		count++
		names = append(names, n)
		if count >= batchSize {
			err = insertFunc(s, names)
			count = 0
			names = names[:0]
			if err != nil {
				return err
			}
		}
	}

	err = insertFunc(s, names[:count])
	if err != nil {
		return err
	}
	return nil
}
