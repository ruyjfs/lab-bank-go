package routes

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	V1(r)
	Graphql(r)
	r.Run()
}
