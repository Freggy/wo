package http

import "wo/storage/blobstore"

type Server struct {
	blobdriver blobstore.Driver
	metadriver blobstore.Driver
}
