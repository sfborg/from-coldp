package sysio

import (
	"os"

	"github.com/gnames/gnsys"
	"github.com/sfborg/from-coldp/internal/ent/sys"
	"github.com/sfborg/from-coldp/pkg/config"
)

type sysio struct {
	cfg config.Config
}

func New(cfg config.Config) sys.Sys {
	return &sysio{cfg: cfg}
}

func (s *sysio) Init() error {
	err := s.cleanup()
	if err != nil {
		return err
	}
	gnsys.MakeDir(s.cfg.CacheSfgaDir)
	return nil
}

func (s *sysio) Close() error {
	return s.cleanup()
}

func (s *sysio) cleanup() error {
	return os.RemoveAll(s.cfg.CacheDir)
}
