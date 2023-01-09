package base

import "webstruct/entity"

type UseCase interface {
	ProcessOperation(entity.Operation) (interface{}, error)
	CreateSession(key string) error
	GetSession(key string) (entity.Session, error)
}
