package static

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/balazskvancz/gorouter"
)

func RegisterStaticRoutes(router gorouter.Router) {
	router.Get("/static/{fileName}", getFile)
}

func getFile(ctx gorouter.Context) {
	files, err := os.ReadDir("uploads")
	if err != nil {
		ctx.SendInternalServerError()
		return
	}

	var (
		paramName = ctx.GetParam("fileName")
		filePath  string
	)

	for _, e := range files {
		if e.IsDir() {
			continue
		}

		if strings.HasPrefix(e.Name(), paramName) {
			filePath = path.Join("uploads", e.Name())
		}
	}

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
