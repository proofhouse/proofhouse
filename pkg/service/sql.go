package service

type SqlService interface {
	Find(id int) bool
}