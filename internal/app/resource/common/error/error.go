package error

import "net/http"

var DBDataInsertFailure = []byte(`{"error": "Database insert failure"}`)
var DBDataReadFailure = []byte(`{"error": "Database read failure"}`)
var ValidationFailure = []byte(`{"error": "Validation failure"}`)
var EmailOrPhoneExists = []byte(`{"error": "Email or phone number already taken"}`)

func InternalServerError(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(error)
}

func ValidationError(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(error)
}

func UnprocessableEntity(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(error)
}
