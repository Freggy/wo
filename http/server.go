package http

import "github.com/freggy/wo/storage/blobstore"

type Server struct {
	blobdriver blobstore.Driver
	metadriver blobstore.Driver
}
