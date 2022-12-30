package database

import (
	"fmt"
)

type ArrayHandler struct {
	ID        int
	Array     string
	Operation string
}

type StructDB struct {
	arrays map[int]ArrayHandler
}

type IStructDBInterface interface {
	AddArray()
	GetArray()
}

func (c *StructDB) AddArray(payload ArrayHandler) error {
	// check if its not the same, if it is, dont do anything
	if _, ok := c.arrays[payload.ID]; ok {
		return fmt.Errorf("no change in '%d' ", payload.ID)
	}
	c.arrays[payload.ID] = payload
	return nil
}

func (c *StructDB) GetArray(payload ArrayHandler) error {
	if _, ok := c.arrays[payload.ID]; ok {
		return nil
	}
	return fmt.Errorf("array not found '%d'", payload.ID)
}
