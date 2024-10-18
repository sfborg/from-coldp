package sfgarc

import "github.com/gnames/coldp/ent/coldp"

type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	InsertMeta(meta *coldp.Meta) error
	Export(outPath string) error
}
