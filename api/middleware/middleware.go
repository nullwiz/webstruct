package middleware

import (
	"crypto/sha256"
	"fmt"
	"webstruct/usecase/base"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware(handlerFunc base.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := handlerFunc.GetSession(createHash(c.ClientIP()))
		// if its an empty interface
		if err != nil {
			{
				new_uuid := createHash((c.ClientIP()))
				// add to context
				c.Set("X-SESSION-ID", new_uuid)
				//c.SetCookie("X-SESSION-ID", new_uuid.String(), 7200, "/", "localhost", false, true)
				// Save session to the database
				err := handlerFunc.CreateSession(new_uuid)
				if err != nil {
					c.JSON(400, gin.H{"error": "Unable to create session"})
				}
				c.Set("X-SESSION-ID", new_uuid)
			}
			//c.SetCookie("X-SESSION-ID", new_uuid.String(), 7200, "/", "localhost", false, true)
			// Save session to the database
		}
		c.Set("X-SESSION-ID", createHash(c.ClientIP()))
		c.Next()
	}
}

func createHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
