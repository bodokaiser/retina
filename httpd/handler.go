package httpd

import (
	"net/http"

	"github.com/bodokaiser/retina/model"
	"github.com/labstack/echo"
)

type Handler struct {
	router *echo.Echo
}

func NewHandler() *Handler {
	h := &Handler{
		router: echo.New(),
	}
	h.router.Post("/", h.Upload)

	return h
}

func (h *Handler) Upload(c *echo.Context) error {
	r := &model.Report{}

	if err := c.Bind(r); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
