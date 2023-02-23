package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madnaaaaas/crud/pkg/refrigerator"
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
			c.JSON(http.StatusOK, "echo")
		})
	}
}

func (r *rest) getBeerByTitle(c *gin.Context) {
	ctx := c.Request.Context()
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, errors.New("empty title"))
		return
	}

	beer, err := r.service.GetBeerByTitle(ctx, title)
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, beerDomainToVM(beer))
}

func (r *rest) postBeer(c *gin.Context) {
	ctx := c.Request.Context()
	vm := new(postBeerVM)
	if err := c.ShouldBindJSON(vm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newBeer, err := postBeerVMToDomain(vm)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := r.service.InsertBeer(ctx, newBeer)
	if err != nil {
		fmt.Println(err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, id)
}
