This repo demonstrates an issue with [Swaggo's](github.com/swaggo/swag) latest generic support, since it's not released we've forked the repo to `github.com/chaintraced/swag/cmd/swag@latest`

Once this version is installed ([Swaggo](github.com/swaggo/swag) from master branch on commit https://github.com/swaggo/swag/commit/9d34a7683d438a2b593fa073ba8820309a3c0987) try to run

```
swag init
```

you'll get an output similar to:

```
2022/08/25 10:40:06 Using overrides from .swaggo
2022/08/25 10:40:06 Generate swagger docs....
2022/08/25 10:40:06 Generate general API Info, search dir:./
2022/08/25 10:40:06 Generating movie.CreateMovie
2022/08/25 10:40:06 Generating $field.Field-Person
2022/08/25 10:40:06 Type definition of type '*ast.IndexExpr' is not supported yet. Using 'object' instead.
2022/08/25 10:40:06 Type definition of type '*ast.IndexExpr' is not supported yet. Using 'object' instead.
2022/08/25 10:40:06 Type definition of type '*ast.IndexExpr' is not supported yet. Using 'object' instead.
2022/08/25 10:40:06 Generating $field.Field-Balance
2022/08/25 10:40:06 create docs.go at  docs/docs.go
2022/08/25 10:40:06 create swagger.json at  docs/swagger.json
2022/08/25 10:40:06 create swagger.yaml at  docs/swagger.yaml[Actor]
```

serve the file with swagger eg.

```
swagger serve ./docs/swagger.json
```
