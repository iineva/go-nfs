package basefs

import "github.com/spf13/afero"

func (fs BaseFS) SymlinkIfPossible(oldname, newname string) error {
	if f, ok := fs.source.(afero.Symlinker); ok {
		return f.SymlinkIfPossible(oldname, newname)
	}
	return ErrNotImplement
}
