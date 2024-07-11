package health

import (
	"net/http"
)

func Get(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("I'm alive!"))
}
