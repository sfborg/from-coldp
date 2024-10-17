package sfgarc

type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	Export(outPath string) error
}
