package httpd

import "net/http"

import "github.com/labstack/echo"
import "github.com/bodokaiser/retina/model"

type Handler struct {
	router *echo.Echo
}

func NewHandler() *Handler {
	h := &Handler{
		router: echo.New(),
	}
	h.router.Get("/images", h.Default)
	h.router.Post("/images", h.Upload)

	return h
}

func (h *Handler) Upload(c *echo.Context) error {
	r := &model.Report{}

	if err := c.Bind(r); err != nil {
		return err
	}
	if err := r.Process(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}

func (h *Handler) Default(c *echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
