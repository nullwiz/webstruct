package base

import "webstruct/entity"

type Repository interface {
	AddSession(entity.Session) (err error)
	GetSession(key string) (entity.Session, error)
	AddOperationToSession(item entity.Operation) (err error)
	GetResultFromSession(entity string, key string) (result interface{}, err error)
	SetResultToSession(entity string, result interface{}, key string) (err error)
	NullResultToSession(entity string, key string) (err error)
}
