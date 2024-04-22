package internal

import (
	"github.com/balazskvancz/gorouter"
	"github.com/balazskvancz/image-gallery-hw/internal/api"
	"github.com/balazskvancz/image-gallery-hw/internal/repository"
	"github.com/balazskvancz/image-gallery-hw/internal/static"
)

// NewServer létrehoz egy új router egyedet és beregisztrálja a végpontokat.
func NewServer(rep repository.Repository) gorouter.Router {
	router := gorouter.New(gorouter.WithAddress(3000))

	api.RegisterRoutes(router, rep)
	static.RegisterStaticRoutes(router)

	return router
}
