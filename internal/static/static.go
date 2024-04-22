package static

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/balazskvancz/gorouter"
)

func RegisterStaticRoutes(router gorouter.Router) {
	router.Get("/static/{fileName}", getFile)
}

func getFile(ctx gorouter.Context) {
	filePath := path.Join("uploads", ctx.GetParam("fileName"))

	f, err := os.Open(filePath)
	if err != nil {
		ctx.SendNotFound()
		return
	}
	defer f.Close()

	imgType := path.Ext(f.Name())[1:]

	ctx.AppendHttpHeader(http.Header{
		"Content-Type": []string{fmt.Sprintf("image/%s", imgType)},
	})

	ctx.Copy(f)
}
