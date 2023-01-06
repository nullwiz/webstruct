package base

import (
	"webstruct/entity"
	stringOp "webstruct/usecase/base/structures/strings"
)

func (ser *Service) ProcessOperation(input entity.Operation) (string, error) {
	var err error
	// get the session
	var session entity.Session
	session, err = ser.Repository.GetSession(input.SessionID)
	if err == nil {
		// Append operation to session
		ser.Repository.AddOperationToSession(input)
		// Process operation
		switch input.Entity {
		case "string":
			switch input.Type {
			case "add":
				ser.Repository.SetResultToSession("String", input.Value, session.ID)
				return "added string", nil
			case "get":
				//get result inteerface
				// cast to string
				// return string
				val, err := ser.Repository.GetResultFromSession("String", session.ID)
				if err == nil && val != "" {
					str, ok := val.(string)
					if ok {
						return str, nil
					}
				}
				return "empty str", err
			case "remove":
				ser.Repository.SetResultToSession("String", "", session.ID)
				return "removed string", nil
			case "removeDups":
				val, err := ser.Repository.GetResultFromSession("String", session.ID)
				// cast val to string
				if err != nil {
					str, ok := val.(string)
					if ok {
						ser.Repository.SetResultToSession("String", stringOp.RemoveDups(str), session.ID)
						return stringOp.RemoveDups(str), nil
					}
				}
			case "palindrome":
				val, err := ser.Repository.GetResultFromSession("String", session.ID)
				// cast val to string
				if err != nil {
					str, ok := val.(string)
					if ok {
						ser.Repository.SetResultToSession("String", stringOp.IsPalindrome(str), session.ID)
						return stringOp.IsPalindrome(str), nil
					}
				}
			}
		default:
			return "not implemented", nil
		}
	}
	return "", err
}

func (ser *Service) CreateSession(key string) error {
	var session entity.Session
	session.ID = key
	session.Operations = make([]*entity.Operation, 0)
	session.Results = make(map[string]interface{})
	ser.Repository.AddSession(session)
	return nil
}

func (ser *Service) GetSession(key string) (entity.Session, error) {
	return ser.Repository.GetSession(key)
}
