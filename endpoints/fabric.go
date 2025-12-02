package endpoints

import (
	"gosql/dsl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeEndpoints(domains []dsl.Domain, router *gin.Engine) {
	for _, domain := range domains {
		for _, endpoint := range domain.Endpoints {
			fullPath := domain.Name + endpoint.RelativePath
			if endpoint.Type == dsl.GET {
				router.GET(fullPath, func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"data": endpoint.Content,
					})
				})
			} else if endpoint.Type == dsl.POST {
				router.POST(fullPath, func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"data": endpoint.Content,
					})
				})
			}
		}
	}
}
