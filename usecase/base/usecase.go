package base

import "webstruct/entity"

type UseCase interface {
	ProcessOperation(entity.Operation) (string, error)
	CreateSession(key string) error
}