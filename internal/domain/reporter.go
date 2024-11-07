package domain

type Reporter interface {
	Generate(result Result)
}
