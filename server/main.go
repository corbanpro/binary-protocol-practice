package main

import (
	"bytes"
	"example/proto"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	requestCountChannel := make(chan bool)

	requestsSent := 0
	router := gin.Default()
	router.POST("/message", func(c *gin.Context) { postMessage(c, requestCountChannel) })
	router.GET("/requests", func(c *gin.Context) { getRequests(c, requestCountChannel, requestsSent) })
	router.GET("/", getIndex)

	go func() {
		lastRequests := requestsSent
		for {
			lastRequests = requestsSent
			time.Sleep(time.Second)
			if requestsSent != lastRequests {
				fmt.Println("Requests per second:", (requestsSent - lastRequests))
			}
		}
	}()

	go func() {
		for range requestCountChannel {
			requestsSent++
		}
	}()

	router.Run("localhost:8080")
}

func getRequests(c *gin.Context, requestCountChannel chan<- bool, requestsSent int) {
	requestCountChannel <- true
	c.JSON(http.StatusOK, gin.H{"requests": requestsSent})
}

func postMessage(c *gin.Context, requestCountChannel chan<- bool) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	bodyBuffer := bytes.NewBuffer(body)

	_, err = proto.Decode(*bodyBuffer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	requestCountChannel <- true

	c.String(http.StatusOK, "OK")
}

type PageData struct {
	Name string
}

func getIndex(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	Index("Corban").Render(c, c.Writer)
}
