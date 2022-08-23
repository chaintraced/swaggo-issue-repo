package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chaintraced/swaggo-issue-repo/movie"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.
		Methods(http.MethodPost).
		Path("/v1/movies/").
		HandlerFunc(
			UploadMove,
		)

	time.Sleep(2 * time.Second)
}

// @title          Example API
// @version        1.0
// @description    Example Description.
// @termsOfService noneyet.com

// @contact.name  API Support
// @contact.url   https://www.chaintraced.com/
// @contact.email nothing@chaintraced.com
// @host          subdomain.chaintraced.com

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
// @description                Add the following to your request header: `Authorization: API-KEY <key>`

// @tag.name        Movie
// @tag.description Movie Public API, these are public api endpoints

// UploadMove godoc
// @tags        Movie
// @Summary     Upload a new Movie
// @Description The API will create a new Movie.
// @Accept      json
// @Produce     json
// @Param       message body     movie.CreateMovie true "Movie Create-Payload"
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /v1/movies/ [post]
// @Security    ApiKeyAuth
func UploadMove(w http.ResponseWriter, r *http.Request) {
	payload := movie.CreateMovie{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
