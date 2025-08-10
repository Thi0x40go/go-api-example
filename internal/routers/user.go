package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Thi0x40go/gin-postgres/internal/models"
	"github.com/Thi0x40go/gin-postgres/internal/views"
)

type ErrorResponse struct {
	Message string `json:message`
}

func RegisterUserRoutes() {
 http.HandleFunc("/users/{id}", getUserHttp)
}

func getUserHttp(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID nao encontrado", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error string", http.StatusBadRequest)
	}

	users := map[int]models.User{
		1: {ID: 1, Name: "Thiago"},
		2: {ID: 2, Name: "Elaine"},
	}

	user, found := users[id]
	if !found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Error ao tentar encontrar o {ID}"})
		return
	}

    response := views.UserResponse(user)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}
