package basefs

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/go-git/go-billy/v5"
	"github.com/spf13/afero"
)

func (fs BaseFS) ReadDir(path string) ([]os.FileInfo, error) {
	f, err := fs.source.Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func (fs BaseFS) Readlink(link string) (string, error) {
	// TODO:
	return "", ErrNotImplement
}

func (fs BaseFS) Lstat(filename string) (os.FileInfo, error) {
	if f, ok := fs.source.(afero.Lstater); ok {
		info, _, err := f.LstatIfPossible(filepath.Clean(filename))
		return info, err
	}
	return nil, ErrNotImplement
}

func (fs BaseFS) CapabilityCheck(capabilities billy.Capability) bool {
	// TODO:
	// fsCaps := Capabilities(fs)
	// return fsCaps&capabilities == capabilities
	return true
}

func (fs BaseFS) Capabilities() billy.Capability {
	return billy.AllCapabilities
}

func (fs BaseFS) Lchown(name string, uid, gid int) error {
	return fs.source.Chown(name, uid, gid)
}

func (fs BaseFS) Symlink(target, link string) error {
	if f, ok := fs.source.(afero.Symlinker); ok {
		return f.SymlinkIfPossible(target, link)
	}
	return ErrNotImplement
}
