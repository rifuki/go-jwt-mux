package productcontroller

import (
	"net/http"

	"github.com/rifuki/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"id":       1,
			"name":     "product a",
			"quantity": 1000,
		},
		{
			"id":       2,
			"name":     "product b",
			"quantity": 1000,
		},
		{
			"id":       3,
			"name":     "product c",
			"quantity": 500,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
