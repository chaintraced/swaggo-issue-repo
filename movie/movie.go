package movie

import "github.com/chaintraced/swaggo-issue-repo/field"

type CreateMovie struct {
	Name           string
	MainActor      field.Field[Person]
	SupportingCast field.Field[[]Person]
	Directors      field.Field[*[]Person]
	CameraPeople   field.Field[[]*Person]
	Costs          field.Field[*Balance]
}

type Person struct {
	Name string
}

type Balance struct {
	Debit  int
	Credit int
}
