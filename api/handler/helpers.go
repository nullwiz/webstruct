package handler

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"webstruct/entity"

	"github.com/gin-gonic/gin"
)

func ErrHandler(err error, c *gin.Context) {
	switch err {
	case entity.ErrEmpty:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	case entity.ErrNotFound:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	case entity.ErrInvalidOperation:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}
}

func Unmarshal(c *gin.Context, t interface{}) error {
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(x, t)

	if err != nil {
		return entity.ErrInvalidInput
	}
	return nil
}

func (input OperationRequest) Validate() error {
	if input.Entity == "" {
		return entity.ErrInvalidOperation
	}
	if input.Type == "" {
		return entity.ErrEmpty
	}
	return nil
}

func createHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
