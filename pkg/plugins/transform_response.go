package plugins

import (
	"github.com/Kong/go-pdk"
)

type TransformResponsePlugin struct{}

func (tr TransformResponsePlugin) Response(kong *pdk.PDK) {
	kong.Response.SetHeader("X-GTW-TESTE", "TESTE")
}

func New() any {
	return &TransformResponsePlugin{}
}
