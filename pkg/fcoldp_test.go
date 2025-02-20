package fcoldp_test

import (
	"path/filepath"
	"testing"

	"github.com/gnames/coldp/ent/coldp"
	fcoldp "github.com/sfborg/from-coldp/pkg"
	"github.com/sfborg/from-coldp/pkg/config"
	"github.com/sfborg/from-coldp/pkg/io/sysio"
	"github.com/sfborg/sflib/io/sfgaio"
	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		msg, path string
	}{
		{"name", "ptero-yaml.zip"},
	}

	for _, v := range tests {
		path := filepath.Join("..", "testdata", "name", v.path)
		arc, err := Init(t, path)
		assert.Nil(err)
		assert.NotNil(arc)
		m, err := arc.Meta()
		assert.Nil(err)
		assert.Equal("CC0", m.License)
	}
}

func Init(t *testing.T, path string) (coldp.Archive, error) {
	assert := assert.New(t)
	cfg := config.New()
	err := sysio.New(cfg).ResetCache()
	assert.Nil(err)

	sfga := sfgaio.New()
	err = sfga.Create(cfg.CacheSfgaDir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	assert.Nil(err)

	fc := fcoldp.New(cfg, sfga)

	return fc.GetCoLDP(path)
}
