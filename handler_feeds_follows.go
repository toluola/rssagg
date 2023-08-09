package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	rssagg "github.com/toluola/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user rssagg.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing Json: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFollowFeeds(r.Context(), rssagg.CreateFollowFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not create feed follow: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedFollowtoFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user rssagg.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not get feed followers: %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedstoFolowers(feedFollow))
}

func (apiCfg *apiConfig) handlerFeedUnfollow(w http.ResponseWriter, r *http.Request, user rssagg.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollwId, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not get Parse feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), rssagg.DeleteFeedFollowParams{
		ID:     feedFollwId,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not delete feed follow: %v", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}
