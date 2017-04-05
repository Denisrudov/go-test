package main

import (
	"names/api"
	"github.com/gin-gonic/gin"
	"sync"
	"net/http"
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 3000, "service port")
	flag.Parse()
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", getHandler)

	r.Run(fmt.Sprintf(":%d", port))

}

func getHandler(c *gin.Context) {
	service := api.NewService()
	var data api.Result
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		randomName, err := service.GetName()
		if err != nil {
			data.Name = api.Name{}
		} else {
			data.Name = randomName
		}

		wg.Done()
	}()

	go func() {
		chukJoke, err := service.GetNorris()
		if err != nil {
			data.ChuckNorris = api.ChuckNorris{}
		} else {
			data.ChuckNorris = chukJoke
		}
		wg.Done()
	}()

	wg.Wait()

	c.JSON(http.StatusOK, data)
}
