package nfs

import (
	"context"
	"net"

	billy "github.com/go-git/go-billy/v5"
	"github.com/willscott/go-nfs/filesystem"
)

// Handler represents the interface of the file system / vfs being exposed over NFS
type Handler interface {
	// Required methods

	Mount(context.Context, net.Conn, MountRequest) (MountStatus, filesystem.FS, []AuthFlavor)

	// Change can return 'nil' if filesystem is read-only
	Change(filesystem.FS) billy.Change

	// Optional methods - generic helpers or trivial implementations can be sufficient depending on use case.

	// Fill in information about a file system's free space.
	FSStat(context.Context, filesystem.FS, *FSStat) error

	// represent file objects as opaque references
	// Can be safely implemented via helpers/cachinghandler.
	ToHandle(fs filesystem.FS, path []string) []byte
	FromHandle(fh []byte) (filesystem.FS, []string, error)
	// How many handles can be safely maintained by the handler.
	HandleLimit() int
}
