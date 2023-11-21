package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gihub.com/prathishbv/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct{
		Name string `json:"name"`
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	ResponseWithJSON(w, 201, databaseUserToUser(user))
}


func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	
	ResponseWithJSON(w, 200, databaseUserToUser(user))
}

