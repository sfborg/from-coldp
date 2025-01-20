package sysio

import (
	"os"

	"github.com/gnames/gnsys"
	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/sfborg/from-coldp/pkg/ent/sys"
)

type sysio struct {
	cfg config.Config
}

func New(cfg config.Config) sys.Sys {
	return &sysio{cfg: cfg}
}

func (s *sysio) ResetCache() error {
	err := os.RemoveAll(s.cfg.CacheDir)
	if err != nil {
		return err
	}
	gnsys.MakeDir(s.cfg.CacheSfgaDir)
	return nil
}
