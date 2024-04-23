package api

import (
	"math/rand"
	"net/http"
	"path"

	"github.com/balazskvancz/gorouter"
	"github.com/balazskvancz/image-gallery-hw/internal/models"
	"github.com/balazskvancz/image-gallery-hw/internal/repository"
)

const (
	nameLength int = 50
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func handleUpload(repo repository.Repository) gorouter.HandlerFunc {
	return func(ctx gorouter.Context) {
		if err := ctx.ParseForm(); err != nil {
			ctx.SendJson(&apiError{Message: "Sikertelen form parseolás!"}, http.StatusBadRequest)
			return
		}

		file, err := ctx.GetFormFile("image")
		if err != nil {
			ctx.SendJson(&apiError{Message: "Sikertelen fájlfeltöltés!"}, http.StatusBadRequest)
			return
		}

		var (
			name      = generateUniqueName(repo)
			extension = path.Ext(file.GetName())
		)

		savePath := path.Join("uploads", name+extension)

		if err := file.SaveTo(savePath); err != nil {
			ctx.SendJson(&apiError{Message: "Sikertelen fájlfeltöltés!"}, http.StatusBadRequest)
			return
		}

		if err := repo.Insert(name); err != nil {
			ctx.SendJson(&apiError{Message: "Sikertelen adatbázismvűvelet!"}, http.StatusBadRequest)
			return
		}

		ctx.SendOk()
	}
}

func handleDelete(repo repository.Repository) gorouter.HandlerFunc {
	return func(ctx gorouter.Context) {
		name := ctx.GetParam("name")
		if name == "" {
			ctx.SendRaw(nil, http.StatusBadRequest, nil)
			return
		}

		if image := repo.GetByName(name); image == nil {
			ctx.SendRaw(nil, http.StatusBadRequest, nil)
			return
		}

		if err := repo.DeleteByName(name); err != nil {
			ctx.SendRaw(nil, http.StatusBadRequest, nil)
			return
		}

		ctx.SendOk()
	}
}

var validDirections []string = []string{"desc", "asc"}

var validOrderBys map[string]string = map[string]string{
	"date": "created_at",
	"name": "name",
}

type GetImagesResponse struct {
	Images models.Images `json:"images"`
}

func handleGet(repo repository.Repository) gorouter.HandlerFunc {
	return func(ctx gorouter.Context) {
		var (
			orderByQuery = ctx.GetQueryParam("orderBy")
			direction    = ctx.GetQueryParam("direction")
		)

		orderBy, contains := validOrderBys[orderByQuery]
		if !contains {
			orderBy = "name"
		}

		if !includes(validDirections, direction) {
			direction = "asc"
		}

		images := repo.GetBy(orderBy, direction)

		ctx.SendJson(&GetImagesResponse{
			Images: images,
		})
	}
}

func generateUniqueName(rep repository.Repository) string {
	buff := make([]rune, nameLength)
	for i := range buff {
		buff[i] = letters[rune(rand.Intn(nameLength))]
	}
	name := string(buff)

	if v := rep.GetByName(name); v != nil {
		return generateUniqueName(rep)
	}

	return name
}
