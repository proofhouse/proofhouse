package service

type Factory interface {
	Create() Service
	Type() string
	Kind() string
}

type Service interface {

}

