package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gihub.com/prathishbv/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Error Parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID: params.FeedId,
	})
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	ResponseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	ResponseWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feeds))
}


func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowstr := chi.URLParam(r, "feedFollowID")
	feedFollowid, err := uuid.Parse(feedFollowstr)
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't Parse feed followID: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowid,
		UserID: user.ID,
	})
	if err != nil {
		ResponseWithErr(w, 400, fmt.Sprintf("Couldn't Delete feed followID: %v", err))
		return
	}
	ResponseWithJSON(w, 201, struct{}{})
}