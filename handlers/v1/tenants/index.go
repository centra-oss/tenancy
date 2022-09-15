package tenants

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.h4n.io/centra/tenancy/models"
	"go.h4n.io/centra/tenancy/models/meta"
)

// @BasePath /api/v1/tenants

// @Summary Get tenants
// @Tags tenants
// @Description Retrieves a list of tenants.
// @Success 200 {object} TenantsIndexResponse
// @Produce json
// @Router / [get]
func TenantsIndex(c *gin.Context) {
    c.JSON(http.StatusOK, TenantsIndexResponse{
        Message: `found 1 tenant(s).`,
        Tenants: []models.Tenant{
            {
                Id: `3e57b452-4c1a-4ec6-b7e2-4a6d712bd8ba`,
                Name: `centra-internal`,
                DisplayName: `Centra Internal`,
                CommonMetadata: meta.CommonMetadata{
                    CreatedAt: time.Now(),
                    UpdatedAt: time.Now(),
                },
            },
        },
    })
}

type TenantsIndexResponse struct {
    Message string `json:"message"`
    Tenants []models.Tenant `json:"tenants"`
}

