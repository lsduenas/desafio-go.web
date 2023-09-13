package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lsduenas/desafio-go-web/cmd/server/handler"
	"github.com/lsduenas/desafio-go-web/internal/domain"
	"github.com/lsduenas/desafio-go-web/internal/tickets"
)

type Router interface {
	MapRoutes()
}
type router struct {
	server  *gin.Engine
	list []domain.Ticket
}

func NewRouter(s *gin.Engine, l []domain.Ticket) Router {
	return &router{
		server:  s,
		list: l,
	}
}
func (router *router) MapRoutes() {
	repo := tickets.NewRepository(router.list)
	service := tickets.NewService(repo)
	handler := handler.NewService(service)
	// group 
	ticketRoutes := router.server.Group("/ticket")
	ticketRoutes.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	ticketRoutes.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
	ticketRoutes.GET("/getAverage/:dest", handler.AverageDestination())
}
