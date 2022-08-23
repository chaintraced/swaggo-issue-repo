This branch demonstrates the fully functional code IF the generic field is within the same package as the parent struct

Here the `.swaggo` file can even beautifully replace the type.

Running `swag init` here give us:

```
2022/08/23 13:40:31 Using overrides from .swaggo
2022/08/23 13:40:31 Generate swagger docs....
2022/08/23 13:40:31 Generate general API Info, search dir:./
2022/08/23 13:40:31 Generating movie.CreateMovie
2022/08/23 13:40:31 Override detected for movie.Field[Actor]: using movie.Actor instead
2022/08/23 13:40:31 Generating movie.Actor
2022/08/23 13:40:31 create docs.go at  docs/docs.go
2022/08/23 13:40:31 create swagger.json at  docs/swagger.json
2022/08/23 13:40:31 create swagger.yaml at  docs/swagger.yaml
```

This is the expected outcome that we'd like on the `main` branch as well.