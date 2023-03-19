package serializers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RequestParams(req *http.Request) map[string]string {
	params := mux.Vars(req)

	return params
}
