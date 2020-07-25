package router

import (
	"github.com/chenrun666/gin_demo/handler/sd"
	"github.com/chenrun666/gin_demo/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}
	return g
}