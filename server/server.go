package server

import (
	"github.com/NavinduNavoda/waggle-gobackend/data"
	"github.com/NavinduNavoda/waggle-gobackend/server/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	dbconfig data.DBConfig
}

func CORSMiddleware_noPkg() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


func NewServer(db data.DBConfig) *Server {

	r := gin.Default()
	r.Use(CORSMiddleware_noPkg())

	api.UserRoutes(r, db)
	api.FilesRoutes(r)


	return &Server{
		router: r,
		dbconfig: db,
	}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}