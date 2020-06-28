package sender

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func New() *Service {
	s := &Service{
		router: gin.Default(),
	}
	s.registerHandlers()
	return s
}

type Service struct {
	router *gin.Engine
}

func (s *Service) Run() error {
	return s.router.Run(":8080")
}

func (s *Service) registerHandlers() {
	s.router.Handle(http.MethodGet, "api/info", func(ctx *gin.Context) {
		type Response struct {
			Number int64 `json:"number"`
		}

		ctx.JSON(http.StatusOK, Response{
			Number: rand.Int63(),
		})
	})
}


