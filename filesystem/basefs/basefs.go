package basefs

import (
	"errors"

	"github.com/spf13/afero"
	"github.com/willscott/go-nfs/filesystem"
)

var (
	ErrNotImplement = errors.New("not implement")
)

type BaseFS struct {
	source afero.Fs
}

func New(source afero.Fs) filesystem.FS {
	return &BaseFS{source: source}
}

func NewMemMapFS() filesystem.FS {
	return New(afero.NewBasePathFs(afero.NewMemMapFs(), "/"))
}

func NewOsFS(path string) filesystem.FS {
	return New(afero.NewBasePathFs(afero.NewOsFs(), path))
}

func NewBasePathFS(source afero.Fs, path string) filesystem.FS {
	return New(afero.NewBasePathFs(source, path))
}
