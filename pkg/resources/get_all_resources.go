package resources

import (
	"net/http"

	"github.com/PsalmsJava/simple-crud-api-with-golong/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllResources(ctx *gin.Context) {
	var resources []models.Resource

	if result := h.DB.Find(&resources); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &resources)
}
