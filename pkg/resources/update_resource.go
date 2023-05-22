package resources

import (
	"net/http"

	"github.com/PsalmsJava/simple-crud-api-with-golong/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateResourceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h handler) UpdateResource(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateResourceRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var resource models.Resource

	if result := h.DB.First(&resource, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	resource.Name = body.Name
	resource.Description = body.Description

	h.DB.Save(&resource)

	ctx.JSON(http.StatusOK, &resource)
}
