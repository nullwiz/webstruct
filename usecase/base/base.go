package base

import (
	"strings"
	"webstruct/entity"
	hashMap "webstruct/usecase/base/structures/hashmap"
	stringOp "webstruct/usecase/base/structures/strings"
)

func (ser *Service) processStringOperation(input entity.Operation, session entity.Session) (string, error) {
	var err error
	var result string
	var sessionID = session.ID
	switch input.Entity {
	case "string":
		switch input.Type {
		case "add":
			ser.Repository.SetResultToSession("String", input.Value, sessionID)
			return "added string", nil
		case "get":
			val, err := ser.Repository.GetResultFromSession("String", sessionID)
			if err == nil && val != "" {
				str, ok := val.(string)
				if ok {
					return str, nil
				}
			}
			return "empty str", err
		case "remove":
			ser.Repository.SetResultToSession("String", "", sessionID)
			return "removed string", nil
		case "removeDups":
			val, err := ser.Repository.GetResultFromSession("String", sessionID)
			if err == nil {
				str, ok := val.(string)
				if ok {
					r := stringOp.RemoveDups(str)
					ser.Repository.SetResultToSession("String", r, session.ID)
					return r, nil
				}
			}
		case "palindrome":
			val, err := ser.Repository.GetResultFromSession("String", sessionID)
			// cast val to string
			if err == nil {
				str, ok := val.(string)
				if ok {
					r := stringOp.IsPalindrome(str)
					ser.Repository.SetResultToSession("String", r, sessionID)
					return r, nil
				}
			}
		}
	default:
		return "not implemented", nil
	}
	return result, err
}
func (ser *Service) processHashMapOperation(input entity.Operation, session entity.Session) (interface{}, error) {
	var err error
	var result string
	var sessionID = session.ID
	switch input.Type {
	case "add":
		ser.Repository.SetResultToSession("HashMap", hashMap.NewHashmap(), sessionID)
		return "added hashmap", nil
	case "get":
		val, err := ser.Repository.GetResultFromSession("HashMap", sessionID)
		if err == nil {
			if err == nil {
				return val, nil
			}
		}
		return "empty hashmap", err
	case "set":
		val, err := ser.Repository.GetResultFromSession("HashMap", sessionID)
		if err == nil {
			hm, ok := val.(*hashMap.Hashmap)
			if ok {
				kv := strings.Split(input.Value, ":")
				k := kv[0]
				v := kv[1]
				hm.Set(k, v)
				ser.Repository.SetResultToSession("HashMap", hm, sessionID)
				r := hm.ReturnCompleteMap()
				if err == nil {
					return r, nil
				}
			}
		}
		return "empty hashmap", err
	case "delete":
		val, err := ser.Repository.GetResultFromSession("HashMap", sessionID)
		if err == nil {
			hm, ok := val.(*hashMap.Hashmap)
			if ok {
				hm.Delete(input.Value)
				ser.Repository.SetResultToSession("HashMap", hm, sessionID)
				return "deleted", nil
			}
		}
	case "clear":
		val, err := ser.Repository.GetResultFromSession("HashMap", sessionID)
		if err == nil {
			hm, ok := val.(*hashMap.Hashmap)
			if ok {
				hm.ClearMap()
				ser.Repository.SetResultToSession("HashMap", hm, sessionID)
				return "cleared", nil
			}
			return "empty hashmap", err
		}
	case "default":
		return "not implemented", nil
	}
	return result, err
}
func (ser *Service) ProcessOperation(input entity.Operation) (interface{}, error) {
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
			return ser.processStringOperation(input, session)
		case "hashmap":
			return ser.processHashMapOperation(input, session)
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
