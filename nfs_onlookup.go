package nfs

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/willscott/go-nfs-client/nfs/xdr"
	"github.com/willscott/go-nfs/filesystem"
)

func lookupSuccessResponse(handle []byte, entPath, dirPath []string, fs filesystem.FS) ([]byte, error) {
	writer := bytes.NewBuffer([]byte{})
	if err := xdr.Write(writer, uint32(NFSStatusOk)); err != nil {
		return nil, err
	}
	if err := xdr.Write(writer, handle); err != nil {
		return nil, err
	}
	if err := WritePostOpAttrs(writer, tryStat(fs, entPath)); err != nil {
		return nil, err
	}
	if err := WritePostOpAttrs(writer, tryStat(fs, dirPath)); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

func onLookup(ctx context.Context, w *response, userHandle Handler) error {
	w.errorFmt = opAttrErrorFormatter
	obj := DirOpArg{}
	err := xdr.Read(w.req.Body, &obj)
	if err != nil {
		return &NFSStatusError{NFSStatusInval, err}
	}

	fs, p, err := userHandle.FromHandle(obj.Handle)
	if err != nil {
		return &NFSStatusError{NFSStatusStale, err}
	}
	contents, err := filesystem.ReadDir(fs, filepath.Join(p...))
	if err != nil {
		return &NFSStatusError{NFSStatusNotDir, err}
	}

	// Special cases for "." and ".."
	if bytes.Equal(obj.Filename, []byte(".")) {
		resp, err := lookupSuccessResponse(obj.Handle, p, p, fs)
		if err != nil {
			return &NFSStatusError{NFSStatusServerFault, err}
		}
		if err := w.Write(resp); err != nil {
			return &NFSStatusError{NFSStatusServerFault, err}
		}
		return nil
	}
	if bytes.Equal(obj.Filename, []byte("..")) {
		if len(p) == 0 {
			return &NFSStatusError{NFSStatusAccess, os.ErrPermission}
		}
		pPath := p[0 : len(p)-1]
		pHandle := userHandle.ToHandle(fs, pPath)
		resp, err := lookupSuccessResponse(pHandle, pPath, p, fs)
		if err != nil {
			return &NFSStatusError{NFSStatusServerFault, err}
		}
		if err := w.Write(resp); err != nil {
			return &NFSStatusError{NFSStatusServerFault, err}
		}
		return nil
	}

	// TODO: use sorting rather than linear
	for _, f := range contents {
		if bytes.Equal([]byte(f.Name()), obj.Filename) {
			newPath := append(p, f.Name())
			newHandle := userHandle.ToHandle(fs, newPath)
			resp, err := lookupSuccessResponse(newHandle, newPath, p, fs)
			if err != nil {
				return &NFSStatusError{NFSStatusServerFault, err}
			}
			if err := w.Write(resp); err != nil {
				return &NFSStatusError{NFSStatusServerFault, err}
			}
			return nil
		}
	}

	fmt.Printf("No file for lookup of %v\n", string(obj.Filename))
	return &NFSStatusError{NFSStatusNoEnt, os.ErrNotExist}
}
