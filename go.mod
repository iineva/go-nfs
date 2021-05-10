module github.com/willscott/go-nfs

go 1.13

require (
	github.com/go-git/go-billy/v5 v5.1.0
	github.com/google/uuid v1.2.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/rasky/go-xdr v0.0.0-20170124162913-1a41d1a06c93
	github.com/spf13/afero v1.6.0
	github.com/willscott/go-nfs-client v0.0.0-20200605172546-271fa9065b33
	github.com/willscott/memphis v0.0.0-20201122065000-f2beb41b6be3
)

// afero not support get uid and gid for now, use this fork to support
replace github.com/spf13/afero => github.com/iineva/afero v1.6.1-0.20210510115905-57c673cfea7b
