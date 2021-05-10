package basefs

import (
	"os"

	"github.com/spf13/afero"
)

func (fs BaseFS) LstatIfPossible(name string) (os.FileInfo, bool, error) {
	if f, ok := fs.source.(afero.Lstater); ok {
		return f.LstatIfPossible(name)
	}
	return nil, false, ErrNotImplement
}
