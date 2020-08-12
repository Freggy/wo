package blobstore

// Tag combines a map digest and a configuration digest. A tag makes it possible to identify
// specific map and configuration contents that are related.
type Tag struct {
	MapDigest    string
	ConfigDigest string
}
