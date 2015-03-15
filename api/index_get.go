package api

import (
	"net/http"
)

type GetIndexHandler struct {
}

func NewGetIndexHandler() *GetIndexHandler {
	return &GetIndexHandler{}
}

func (g *GetIndexHandler) name() string {
	return "test"
}

func (g *GetIndexHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {

}
