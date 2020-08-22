package main

import (
	"errors"
	"github.com/chenrun666/gin_demo/config"
	"github.com/chenrun666/gin_demo/model"
	"github.com/chenrun666/gin_demo/router"
	"github.com/chenrun666/gin_demo/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	//var middlewares []gin.HandlerFunc
	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.Logging(),
	}

	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address:%s\n", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp != nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect to the router")
}
