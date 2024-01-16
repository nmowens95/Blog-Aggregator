package main

import (
	"fmt"
	"net/http"

	"http.github.com/nmowens95/Blog-Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %s", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
