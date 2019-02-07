package swaggerui

import (
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"net/http"
	"strings"
)

func Handler(prefix, jsonUrl string) http.Handler {
	box := packr.NewBox("./slim")
	fs := &indexInserter{FileSystem: box, jsonUrl: jsonUrl}
	return http.StripPrefix(prefix, http.FileServer(fs))
}

type indexInserter struct {
	http.FileSystem
	jsonUrl string
}

func (i *indexInserter) Open(name string) (http.File, error) {
	file, err := i.FileSystem.Open(name)

	if name == "/index.html" {
		b, _ := ioutil.ReadAll(file)
		str := strings.Replace(string(b), "https://petstore.swagger.io/v2/swagger.json", i.jsonUrl, -1)
		reader := strings.NewReader(str)
		return packd.NewFile(name, reader)
	}

	return file, err
}
