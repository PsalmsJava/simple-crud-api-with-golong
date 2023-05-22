package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PsalmsJava/simple-crud-api-with-golong/pkg/common/models"
)

func (h handler) DeleteResource(ctx *gin.Context) {
	id := ctx.Param("id")

	var resource models.Resource

	if result := h.DB.First(&resource, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&resource)

	ctx.Status(http.StatusOK)
}
