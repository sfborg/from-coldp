package fcoldp

import (
	"log/slog"
	"os"
	"sync"

	"github.com/gnames/coldp/ent/coldp"
)

func (fc *fcoldp) importData() error {
	paths := fc.c.DataPaths()
	for k, v := range paths {
		switch k {
		case coldp.NameDT:
			_ = fc.importName(v)
		}
	}
	return nil
}

func (fc *fcoldp) importName(path string) error {
	chIn := make(chan coldp.Name)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for n := range chIn {
			err := fc.s.InsertName(n)
			if err != nil {
				slog.Error("AAAA", "error", err)
				os.Exit(1)
			}
		}
	}()

	err := coldp.Read(fc.c.Config(), path, chIn)
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}
