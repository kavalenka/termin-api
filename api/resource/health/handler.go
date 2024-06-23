package health

import (
	"fmt"
	"net/http"
	"termin-api/config"
)

func Get(w http.ResponseWriter, _ *http.Request) {
	dbConfig := config.NewDB()
	dbConfigStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	w.Write([]byte(dbConfigStr))
}
