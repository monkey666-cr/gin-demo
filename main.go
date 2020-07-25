package main

import (
	"errors"
	"github.com/chenrun666/gin_demo/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	g := gin.New()

	var middlewares []gin.HandlerFunc

	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatalln("The router has no response, or it might took too long to start up.", err)
		}
		log.Println("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address:%s\n", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}


func pingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp != nil && resp.StatusCode == 200{
			return nil
		}

		log.Println("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to the router")
}