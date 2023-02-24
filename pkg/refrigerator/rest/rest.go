package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madnaaaaas/crud/pkg/refrigerator"
	"github.com/madnaaaaas/crud/pkg/utils"
)

type rest struct {
	service refrigerator.Service
}

func NewRest(service refrigerator.Service) *rest {
	return &rest{service}
}

func (r *rest) Register(api *gin.RouterGroup) {
	route := api.Group("/refrigerator")
	{
		route.GET(":title", r.getBeerByTitle)
		route.POST("", r.postBeer)
		route.GET("echo", func(c *gin.Context) {
			utils.PublishData(c, "echo")
		})
	}
}

func (r *rest) getBeerByTitle(c *gin.Context) {
	ctx := c.Request.Context()
	title := c.Param("title")
	if title == "" {
		utils.PublishError(c, errors.New("empty title"), http.StatusBadRequest)
		return
	}

	beer, err := r.service.GetBeerByTitle(ctx, title)
	if err != nil {
		utils.PublishError(c, err, http.StatusInternalServerError)
		return
	}

	utils.PublishData(c, beerDomainToVM(beer))
}

func (r *rest) postBeer(c *gin.Context) {
	ctx := c.Request.Context()
	vm := new(postBeerVM)
	if err := c.ShouldBindJSON(vm); err != nil {
		utils.PublishError(c, err, http.StatusBadRequest)
		return
	}

	newBeer, err := postBeerVMToDomain(vm)
	if err != nil {
		utils.PublishError(c, err, http.StatusBadRequest)
		return
	}

	id, err := r.service.InsertBeer(ctx, newBeer)
	if err != nil {
		utils.PublishError(c, err, http.StatusInternalServerError)
		return
	}

	utils.PublishData(c, id)
}
