package handler

import (
	"webstruct/api/middleware"
	"webstruct/entity"
	"webstruct/usecase/base"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	Usecase base.UseCase
}

func NewGinHandler(usecase base.UseCase) *gin.Engine {
	h := &GinHandler{Usecase: usecase}
	r := gin.Default()
	r.Use(middleware.SessionMiddleware(h.Usecase))
	r.POST("/structures/operation", h.ProcessOperation)
	return r
}

func (h *GinHandler) ProcessOperation(c *gin.Context) {
	var l OperationRequest

	err := Unmarshal(c, &l)
	if err != nil {
		ErrHandler(err, c)
		return
	}
	err = l.Validate()
	if err != nil {
		ErrHandler(err, c)
		return
	}
	var op entity.Operation
	op.Entity = l.Entity
	op.Type = l.Type
	if l.Value != "" {
		op.Value = l.Value
	}
	//op.SessionID, err = c.Cookie("X-SESSION-ID")
	// get from context
	op.SessionID = c.GetString("X-SESSION-ID")
	if err != nil {
		ErrHandler(err, c)
		return
	}

	result, err := h.Usecase.ProcessOperation(op)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res := OperationResponse{Result: result}
	c.JSON(200, res)
}
