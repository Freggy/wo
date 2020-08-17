### Authentication

For now only basic auth is possible. In the future it will be possible to use JWT tokens for authentication.

Provide the `Authorization: Basic <credentials>` header in requests made to the registry. 
If the credentials are invalid the server will respond with `401 Unauthorized`.

### Errors

Every response with status code within `4xx` or `5xx` will contain an error object to provide error related information.

```JSON
{
  "msg": "this is an error"
}
```

### Get version of the registry

```
GET /<major_version>
```

Returns 
* `200 OK` if the registry implements `major_version`
* `404 Not Found` if the registry does not implement `major_version`

### Upload a map blob

```
PUT /blobs/upload
Content-Type: application/octet-stream

<map binary data>
```

Uploads a map blob to the registry.

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`.
The response will always provide the map digest in the `Wo-Map-Digest`.

```
201 Created
Wo-Map-Digest: <digest>
```

### Upload a config file

```
PUT /configs/upload
Content-Type: text/plain

<config text data>
```

Uploads a config to the registry.

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`.
The response will always provide the map digest in the `Wo-Config-Digest`.

```
201 Created
Wo-Config-Digest: <digest>
```

### Create a tag

```
PUT /<repository>/<map>/<tag_id>
Content-Type: application/json

{
  "config_digest": "<digest>",
  "map_digest": "<digest>"
}
```

Creates a tag referencing the given map and config digest.

If the upload succeeded the response will contain a `201 Created`. When the blob already exists the response will contain `409 Conflict`. 
`404 Not Found` is returned when either `config_digest` or `map_digest` point to an invalid or non existing reference.

### Set/update map metadata

```
PUT /<repository>/<map>/metadata
Content-Type: text/plain

key1=value1 key2=value2
...
```

This endpoint can be used to set or update custom metadata about a specific map. Only simple key value pairs are possible.

Successful request will return `200 OK`. If either `<repository>` or `<map>` could not be found `404 Not Found` is returned.

### Get map metadata

```
GET /<repository>/<map>/metadata/<key>
```

Retrieves the value of a given key of map specific metadata.

Return `200 OK` if key can be found. `404 Not Found` when key does not exist.

```
200 OK
Content-Type: text/plain

<value>
```

### Get map by tag

```
GET /<repository>/<map>/<tag_id>
```

All the content referenced by this tag will be returned in a gzipped tarball in the form of

```
├── map
├── metadata
└── config
```

The client will receive a `404 Not Found` if the specified tag does not exist.

```
200 OK

<gzipped tarball>
```

### Delete map metadata

```
DELETE /<repository>/<map>/metadata/<key>
```

Delete map specific metadata

Returns `200 OK` if successful. 

