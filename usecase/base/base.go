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
	if err != nil {
		// Append operation to session
		ser.Repository.AddOperationToSession(input)
		// Process operation
		switch input.Entity {
		case "string":
			switch input.Type {
			case "add":
				ser.Repository.SetResultToSession("String", input.Value, session.ID)
			case "remove":
				ser.Repository.SetResultToSession("String", "", session.ID)
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
			return "", nil
		}
	}
	return "", nil
}

func (ser *Service) CreateSession(key string) error {
	var session entity.Session
	session.ID = key
	session.Operations = make([]*entity.Operation, 0)
	session.Results = make(map[string]interface{})
	ser.Repository.AddSession(session)
	return nil
}
