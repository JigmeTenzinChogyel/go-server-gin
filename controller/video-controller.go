package controller

import (
	"go-server/entity"
	"go-server/service"
	"go-server/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	vidoes := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": vidoes,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}
