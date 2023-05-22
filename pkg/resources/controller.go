package resources

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/books")
	routes.POST("/", h.AddResource)
	routes.GET("/", h.GetAllResources)
	routes.GET("/:id", h.GetSingleResource)
	routes.PUT("/:id", h.UpdateResource)
	routes.DELETE("/:id", h.DeleteResource)
}
