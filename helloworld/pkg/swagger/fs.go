package swagger

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
)

// BindataFS go-bindata 命令使用方式
func BindataFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "third_party/swagger-ui",
	}
}

// AssetFS go-bindata-assetfs 使用方式
func AssetFS() *assetfs.AssetFS {
	return assetFS()
}
