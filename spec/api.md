### Authentication

`TODO`

### Errors

`TODO`

### Upload a map blob

Uploads a map blob to the registry.

```
PUT /blobs/upload
Content-Type: application/octet-stream

<map binary data>
```

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`.
The response will always provide the map digest in the `Wo-Map-Digest`.

```
201 Created
Wo-Map-Digest: <digest>
```

### Upload a config file

Uploads a config to the registry.

```
PUT /configs/upload
Content-Type: text/plain

<config text data>
```

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`.
The response will always provide the map digest in the `Wo-Config-Digest`.

```
201 Created
Wo-Config-Digest: <digest>
```

### Create a tag

Creates a tag referencing the given map and config digest.

```
PUT /<repository>/<map>/<tag_id>
Content-Type: application/json

{
  "config_digest": "<digest>",
  "map_digest": "<digest>"
}
```

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`. 
`404 Not Found` is returned when either `config_digest` or `map_digest` point to an invalid or non existing reference.

### Set/update map metadata

This endpoint can be used to set or update custom metadata about a specific map. Only simple key value pairs are possible. Nested keys will be ignored.

```
PUT /<repository>/<map>/metadata
Content-Type: text/plain

key1=value1 key2=value2
...
```

Successful request will return `200 OK`. If either `<repository>` or `<map>` could not be found `404 Not Found` is returned.

### Get map metadata

Retrieves the value of a given key of map specific metadata.

```
GET /<repository>/<map>/metadata/<key>
```

Return `200 OK` if key can be found. `404 Not Found` when key does not exist.

```
200 OK
Content-Type: text/plain

<value>
```

### Get map by tag

All the content referenced by this tag will be returned in a gzipped tarball in the form of

```
├── map
├── metadata
└── config
```

```
GET /<repository>/<map>/<tag_id>
```

The client will receive a `404 Not Found` if the specified tag does not exist.

```
200 OK

<gzipped tarball>
```

### Delete map metadata

Delete map specific metadata

```
DELETE /<repository>/<map>/metadata/<key>
```

Returns `200 OK` if successful. 



