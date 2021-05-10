package helpers

import (
	"context"
	"log"
	"net"

	"github.com/go-git/go-billy/v5"
	"github.com/willscott/go-nfs"
	"github.com/willscott/go-nfs/filesystem"
	"github.com/willscott/go-nfs/filesystem/basefs"
)

const (
	defaultCreateMode = 0755
)

// NewNullAuthHandler creates a handler for the provided filesystem
func NewNullAuthHandler(fs filesystem.FS) nfs.Handler {
	return &NullAuthHandler{fs}
}

// NullAuthHandler returns a NFS backing that exposes a given file system in response to all mount requests.
type NullAuthHandler struct {
	fs filesystem.FS
}

// Mount backs Mount RPC Requests, allowing for access control policies.
func (h *NullAuthHandler) Mount(ctx context.Context, conn net.Conn, req nfs.MountRequest) (status nfs.MountStatus, hndl filesystem.FS, auths []nfs.AuthFlavor) {
	status = nfs.MountStatusOk
	if err := h.fs.MkdirAll(string(req.Dirpath), defaultCreateMode); err != nil {
		// TODO:
		log.Printf("mout create dir error on: %v", err.Error())
	}
	hndl = basefs.NewBasePathFS(h.fs, string(req.Dirpath))
	// hndl = h.fs
	auths = []nfs.AuthFlavor{nfs.AuthFlavorNull}
	return
}

// Change provides an interface for updating file attributes.
func (h *NullAuthHandler) Change(fs filesystem.FS) billy.Change {
	if c, ok := fs.(billy.Change); ok {
		return c
	}
	return nil
}

// FSStat provides information about a filesystem.
func (h *NullAuthHandler) FSStat(ctx context.Context, f filesystem.FS, s *nfs.FSStat) error {
	return nil
}

// ToHandle handled by CachingHandler
func (h *NullAuthHandler) ToHandle(f filesystem.FS, s []string) []byte {
	return []byte{}
}

// FromHandle handled by CachingHandler
func (h *NullAuthHandler) FromHandle([]byte) (filesystem.FS, []string, error) {
	return nil, []string{}, nil
}

// HandleLImit handled by cachingHandler
func (h *NullAuthHandler) HandleLimit() int {
	return -1
}
