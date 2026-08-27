package mux

import (
	"errors"
	"net/http"
)

// GetParamValue to get param value
func GetParamValue(r *http.Request, param string, variable string) (string, error) {
	if len(variable) == 0 {
		variable = r.FormValue(param)
		if len(variable) == 0 {
			urlParam := r.URL.Query()
			variable = urlParam.Get(param)
			if len(variable) == 0 {
				return "", errors.New("Bad request")
			}
		}
	}

	return variable, nil
}
