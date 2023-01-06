package mockdatabase

import (
	"strconv"
	"time"
	"webstruct/entity"
)

type MockDatabase struct {
	// map of operation entities
	// key is the operation ID
	// value is the operation entity
	sessions   map[string]entity.Session
	operations map[string]entity.Operation
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		sessions:   make(map[string]entity.Session),
		operations: make(map[string]entity.Operation),
	}
}
func (r *MockDatabase) AddSession(item entity.Session) (err error) {
	r.sessions[item.ID] = item
	return nil
}
func (r *MockDatabase) GetSession(key string) (entity.Session, error) {
	// if there is a session, return it, if not, fail
	if _, ok := r.sessions[key]; ok {
		return r.sessions[key], nil
	}
	return entity.Session{}, entity.ErrSessionNotFound
}
func (r *MockDatabase) RecordTimestampOnSession(key string) (err error) {
	var session = r.sessions[key]
	// add time stamp
	session.Timestamp = time.Now().Unix()
	return nil
}

func (r *MockDatabase) AddOperationToSession(item entity.Operation) (err error) {
	// autoincrement id
	item.ID = strconv.Itoa(len(r.operations) + 1)
	var session = r.sessions[item.SessionID]
	session.Operations = append(r.sessions[item.SessionID].Operations, &item)
	r.sessions[item.SessionID] = session
	return nil
}
func (r *MockDatabase) GetResultFromSession(entity string, key string) (result interface{}, err error) {
	var session = r.sessions[key]
	return session.Results[entity], nil
}

func (r *MockDatabase) ClearOperationsFromSession(key string) (err error) {
	var session = r.sessions[key]
	session.ClearOperations()
	return nil
}
func (r *MockDatabase) SetResultToSession(entity string, result interface{}, key string) (err error) {
	var session = r.sessions[key]
	// add result to session key
	session.SetResult(entity, result)

	return nil
}

func (r *MockDatabase) NullResultToSession(entity string, key string) (err error) {
	var session = r.sessions[key]
	session.NullResult(entity)
	return nil
}

func (r *MockDatabase) ClearResultsFromSession(key string) (err error) {
	var session = r.sessions[key]
	session.Results = map[string]interface{}{}
	return nil
}
