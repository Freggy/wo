Storage
=======

Maps and configs are content-addressable. This means they can be identified by their content rather. They are saved to the file system in the following format:

```
<digest of content>
└── <content>
```

`digest` is defined as the hexadecimal string representation of the SHA-256 hash of the content being addressed.  

```
digest := sha256(content)
```

All following paths are located under the default base path `/var/lib/wo`. Note that the base path can be configured differently.

Maps
----

Each map blob is stored under `/maps/<digest>`.

```
<digest of content>
└── map

maps
├── 4d9a3555e52bcfca2844438ff707ba0dd029df47ab5b978be5019287c6c97999
│   └── map
└── ...
```

`map` is an arbitrary binary large object. The registry does not know the exact type of content `map` is. 

Configs
-------

Each configuration file is stored under `/configs/<digest>`.

```
<digest of content>
└── config

configs
├── 6ca13d52ca70c883e0f0bb101e425a89e8624de51db2d2392593af6a84118090
│   └── config
└── ...
```

Tags
----

Each tag is stored under `/tags/<repository>/<map_name>`. Tags are plain JSON files.

```
tags
└── my-repository
    └── my-map
        ├── 1.0.0.json
        ├── 2.0.0.json
        └── ...
```

Metadata
--------

Metadata for each map is stored under `/metadata/<repository>/<map_name>`.

The content of the metadata file is in plain ini-like style:

```
<key>=<value>
<key>=<value>
<key>=<value>
```

Each line represents one key value pair.

```
metadata
└── my-repository
    └── my-map
        └── metadata
```


