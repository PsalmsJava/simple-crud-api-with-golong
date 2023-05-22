package resources

import (
	"net/http"

	"github.com/PsalmsJava/simple-crud-api-with-golong/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddResourceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h handler) AddResource(ctx *gin.Context) {
	body := AddResourceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var resource models.Resource

	resource.Name = body.Name
	resource.Description = body.Description

	if result := h.DB.Create(&resource); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &resource)
}
