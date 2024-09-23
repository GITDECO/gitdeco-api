package delivery

import (
	"gitdeco-api/internal/exception"
	"gitdeco-api/internal/svg"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SvgHandler struct {
	svgUsecase svg.Usecase
}

func NewSvgHandler(svgUsecase svg.Usecase) svg.Handler {
	return &SvgHandler{svgUsecase: svgUsecase}
}

func (dh *SvgHandler) GetBadge(c *fiber.Ctx) error {
	text := c.Query("text")
	svg, err := dh.svgUsecase.GetBadge(text)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	c.Set("Content-Type", "image/svg+xml")
	return c.Status(http.StatusOK).Send(svg)
}

func (dh *SvgHandler) GetTemplate(c *fiber.Ctx) error {
	headerName := c.Query("header-name")
	title := c.Query("title")
	projectTitle := c.Query("project-title")
	projectSubTitle := c.Query("project-sub-title")
	projectDescription1 := c.Query("project-description1")
	projectDescription2 := c.Query("project-description2")
	svg, err := dh.svgUsecase.GetTemplate(headerName, title, projectTitle, projectSubTitle, projectDescription1, projectDescription2)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	c.Set("Content-Type", "image/svg+xml")
	return c.Status(http.StatusOK).Send(svg)
}

func (dh *SvgHandler) GetTemplate2(c *fiber.Ctx) error {
	headerName := c.Query("header-name")
	title := c.Query("title")
	projectTitle := c.Query("project-title")
	projectSubTitle := c.Query("project-sub-title")
	projectDescription := c.Query("project-description")
	backgroundImage := c.Query("background-image")
	profileImage := c.Query("profile-image")
	projectIntroImage := c.Query("project-intro-image")
	techStackList := c.Query("tech-stack-list")
	svg, err := dh.svgUsecase.GetTemplate2(headerName, title, projectTitle, projectSubTitle, projectDescription, backgroundImage, profileImage, projectIntroImage, techStackList)
	if err != nil {
		panic(&exception.Error{Key: "USECASE_ERROR", Data: err.Error()})
	}
	c.Set("Content-Type", "image/svg+xml")
	return c.Status(http.StatusOK).Send(svg)
}