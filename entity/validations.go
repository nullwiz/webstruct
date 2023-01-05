package entity

import "errors"

func (input Session) Validate() error {
	if input.ID == "" {
		return errors.New("empty ID")
	}
	return nil
}

func (input Operation) Validate() error {
	if input.SessionID == "" {
		return errors.New("empty ID")
	}

	if input.Entity != "String" {
		return errors.New("invalid entity for operation")
	}
	if input.Type == "" {
		return errors.New("empty type")
	}
	// if the type matches an entity and its

	return nil
}
