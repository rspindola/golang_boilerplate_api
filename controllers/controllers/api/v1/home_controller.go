package v1

import (
	"net/http"

	"github.com/rspindola/boilerplate_golang_api/helpers/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
