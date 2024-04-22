package api

import (
	"net/http"

	"github.com/balazskvancz/gorouter"
	"github.com/balazskvancz/image-gallery-hw/internal/repository"
)

const (
	routeInsert string = "/api/images"
	routeDelete string = "/api/images/{name}"
	routeGet    string = "/api/images"
)

var (
	allowedResponseHeaders string = "Access-Control-Allow-Origin, Authorization, Content-Type"
	allowedMethods         string = http.MethodGet + ", " + http.MethodPost + ", " + http.MethodDelete + ", " + http.MethodPut
)

func cors() gorouter.Middleware {
	var mw = func(ctx gorouter.Context, next gorouter.HandlerFunc) {
		corsHeader := http.Header{
			"Access-Control-Allow-Origin":      []string{"http://localhost:8080"},
			"Access-Control-Allow-Headers":     []string{allowedResponseHeaders},
			"Access-Control-Allow-Methods":     []string{allowedMethods},
			"Access-Control-Allow-Credentials": []string{"true"},
		}

		ctx.AppendHttpHeader(corsHeader)

		next(ctx)

	}

	return gorouter.NewMiddleware(mw)
}

// RegisterRoutes beregisztrálja az API végpontokat.
func RegisterRoutes(router gorouter.Router, repo repository.Repository) {
	router.RegisterMiddlewares(cors())

	router.Post(routeInsert, handleUpload(repo))
	router.Delete(routeDelete, handleDelete(repo))
	router.Get(routeGet, handleGet(repo))
}