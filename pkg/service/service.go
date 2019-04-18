package service

type Service interface {
	Type() string
	Kind() string
}
