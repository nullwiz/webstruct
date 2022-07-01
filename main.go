package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	Arrays "webstruct/arrays"
)
//import functions from array file 

var objectStore = ""
var arrayOps = []string{"add", "remove", "get", "remDups", "toLowcase"}

func init(){
	// Something to do here
}

func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  router.Use(gin.Logger())
  log.Println("Starting server...")

  router.Use(static.Serve("/", static.LocalFile("./views", true)))
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }
  api.GET("/structures", StructHandler)
  api.POST("/structures/arrays", ArraysHandlerPost)
  router.Run(":3000")
}

func validateOp(op string) bool{
	for _, opVal := range arrayOps {
		if opVal == op {
			return true
		}
	}
	return false
}

// Lists all structures 
func StructHandler(c *gin.Context){
	c.Header("Content-Type", "application/json")
	files, err:= ioutil.ReadDir("./structures")
	
	if err!=nil{
		log.Fatal(err)
	}
	structs := []string{}
	for _, file:= range files {
		structs = append(structs, file.Name())
	}	
	c.JSON(http.StatusOK, gin.H{ "message" : structs })
}

func ArraysHandlerPost(c *gin.Context){
	c.Header("Content-Type", "application/json")
	op := c.Query("op")
	if !validateOp(op){
		c.JSON(http.StatusBadRequest, gin.H{ "message" : "Invalid operation"})
		return
	}
	switch op {
		case "add":
			array := c.Query("array")
			objectStore = array
			c.JSON(http.StatusOK, gin.H{ "message" : "Array added to objectstore: " + objectStore})
		case "remove":
			objectStore = ""
			c.JSON(http.StatusOK, gin.H{ "message" : "Array removed from objectstore"})
		case "get":
			c.JSON(http.StatusOK, gin.H{ "message" : objectStore})
		case "remDups":
			objectStore = Arrays.RemoveDuplicatesFromStringLiteral(objectStore)
			c.JSON(http.StatusOK, gin.H{ "message" : objectStore})
		case "toLowcase":
			objectStore = Arrays.Lowercase(objectStore)
			c.JSON(http.StatusOK, gin.H{ "message" : objectStore})
		default:
			c.JSON(http.StatusBadRequest, gin.H{ "message" : "Invalid operation"})
			return
	}
}
