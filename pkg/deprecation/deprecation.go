package deprecation

import (
	"errors"
	"net/http"
)

var ErrDeprecation = errors.New("server indicated that this client is deprecated")

func CheckResponseDeprecation(r *http.Response) error {
	if r.Header.Get("Deprecated-Client") == "true" {
		return ErrDeprecation
	}

	return nil
}
