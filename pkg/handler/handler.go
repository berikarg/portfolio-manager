package handler

import (
	"github.com/berikarg/portfolio-manager/pkg/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		assets := api.Group("/assets")
		{
			assets.POST("/", h.createAsset)
			assets.GET("/", h.getAllAssets)
			assets.GET("/:id", h.getAssetById)
			assets.PUT("/:id", h.updateAsset)
			assets.DELETE("/:id", h.deleteAsset)

			savings := assets.Group(":id/savings")
			{
				savings.POST("/", h.createSaving)
				savings.GET("/", h.getAllSavings)
			}
		}

		savings := api.Group("savings")
		{
			savings.GET("/:id", h.getSavingById)
			savings.PUT("/:id", h.updateSaving)
			savings.DELETE("/:id", h.deleteSaving)
		}
	}
	return router
}
