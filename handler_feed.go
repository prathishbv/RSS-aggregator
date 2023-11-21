package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gihub.com/prathishbv/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	ResponseWithJSON(w, 201, databaseFeedToFeed(feed))
}


func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	ResponseWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}


