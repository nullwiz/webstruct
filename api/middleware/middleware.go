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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
func createHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
