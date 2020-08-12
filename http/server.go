package http

import (
	"github.com/freggy/wo/storage/blobstore/driver"
)

type Server struct {
	blobdriver driver.Driver
	metadriver driver.Driver
}
