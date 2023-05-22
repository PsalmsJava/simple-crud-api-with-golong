package resources

import (
	"net/http"

	"github.com/PsalmsJava/simple-crud-api-with-golong/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetSingleResource(ctx *gin.Context) {
	id := ctx.Param("id")

	var resource models.Resource

	if result := h.DB.First(&resource, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &resource)
}
