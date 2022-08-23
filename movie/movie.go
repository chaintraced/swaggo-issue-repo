package movie

import "github.com/chaintraced/swaggo-issue-repo/field"

type CreateMovie struct {
	Name  string
	Actor field.Field[Actor]
}

type Actor struct {
	Name string
}
