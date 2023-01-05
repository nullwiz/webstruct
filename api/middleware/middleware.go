package middleware

import (
	"webstruct/usecase/base"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func SessionMiddleware(handlerFunc base.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := c.GetHeader("X-SESSION-ID")
		if sessionId == "" {
			// generate a uuid based session id based on the ip address
			// and store it in the header
			ip := c.ClientIP()
			new_uuid := uuid.NewV5(uuid.NamespaceOID, ip)
			// Save session to the database
			err := handlerFunc.CreateSession(new_uuid.String())
			if err != nil {
				c.JSON(400, gin.H{"error": "Unable to create session"})
			}
			c.Set("X-SESSION-ID", new_uuid.String())
		}
		c.Next()
	}
}
