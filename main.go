package main

import (
	"io/ioutil"
	"log"
	"net/http"
	array "webstruct/arrays"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var objectStore = ""
var arrayOps = []string{"add", "remove", "get", "remDups", "toLowcase", "isPalindrome"}
var hashOps = []string{"addKey", "removeKeyValue", "get", "remDups"}
var arrayHandler = new(array.ArrayH)

func init() {
	// Something to do here
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(gin.Logger(), CORSMiddleware())
	log.Println("Starting server...")

	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	api.GET("/structures", StructHandler)
	api.POST("/structures/arrays", ArraysHandlerPost)
	router.Run(":3000")
}

// Lists all structures
func StructHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	files, err := ioutil.ReadDir("./structures")

	if err != nil {
		log.Fatal(err)
	}
	structs := []string{}
	for _, file := range files {
		structs = append(structs, file.Name())
	}
	c.JSON(http.StatusOK, gin.H{"message": structs})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ArraysHandlerPost(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	op := c.Query("op")
	array := c.DefaultQuery("array", "NOP")
	newObject := arrayHandler.handler{ID: 1, Array: "sarasa", Operation: "NOP"}
}

// func HashMapHandlerPost(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	op := c.Query("op")
// 	if !validateOp(op) {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid operation"})
// 		return
// 	}
// 	switch op {
// 	case "addKey":
// 		hashMap := c.Query("hashmap")
// 		objectStore = hashMap
// 		c.JSON(http.StatusOK, gin.H{"message": "Array added", "array": objectStore})
// 	case "remove":
// 		objectStore = ""
// 		c.JSON(http.StatusOK, gin.H{"message": "Array removed"})
// 	case "get":
// 		c.JSON(http.StatusOK, gin.H{"message": "Array retrieved", "array": objectStore})
// 	case "remDups":
// 		objectStore = HashMap.RemoveDuplicatesFromStringLiteral(objectStore)
// 		c.JSON(http.StatusOK, gin.H{"message": "Array removed duplicates", "array": objectStore})
// 	case "toLowcase":
// 		objectStore = HashMap.Lowercase(objectStore)
// 		c.JSON(http.StatusOK, gin.H{"message": "Array to lowercase", "array": objectStore})
// 	case "isPalindrome":
// 		if HashMap.IsPalindrome(objectStore) {
// 			c.JSON(http.StatusOK, gin.H{"message": "true", "array": objectStore})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"message": "false", "array": objectStore})
// 		}
// 	default:
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid operation"})
// 		return
// 	}
// }
