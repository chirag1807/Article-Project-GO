package controller

import (
	testcase "articleproject/api/model/dto/test-cases"
	"articleproject/api/model/response"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	r.Post("/api/user/follow/:ID", NewFollowerController(followerService).Follow)

	req, _ := http.NewRequest("POST", "/api/user/follow/:ID", http.NoBody)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("ID", "941445337509003265")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	var followersid int64 = 941150051613179905
	ctx = context.WithValue(ctx, ContextKeyID, followersid)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var response testcase.Follower
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "You are Following the Author Now.", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestUnFollow(t *testing.T) {
	r.Delete("/api/user/unfollow/:ID", NewFollowerController(followerService).UnFollow)

	req, _ := http.NewRequest("DELETE", "/api/user/unfollow/:ID", http.NoBody)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("ID", "941445337509003265")
	ctx := context.WithValue(req.Context(), chi.RouteCtxKey, rctx)
	var followersid int64 = 941150051613179905
	ctx = context.WithValue(ctx, ContextKeyID, followersid)
	req = req.WithContext(ctx)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var response testcase.Follower
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "You are Not Following the Author Now.", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestMyFollowers(t *testing.T) {
	expected := []response.Follower{
		{
			ID:    941445337509003265,
			Name:  "Sameer Makwana",
			Bio:   "Junior Software Engineer at ZURU TECH INDIA.",
			Image: nil,
		},
		{
			ID:    942546435960995841,
			Name:  "Dhyey Panchal",
			Bio:   "Junior Software Engineer at Rapidops INC.",
			Image: nil,
		},
	}
	r.Get("/api/user/myfollowers", NewFollowerController(followerService).MyFollowers)

	req, _ := http.NewRequest("GET", "/api/user/myfollowers", http.NoBody)

	var id int64 = 941150051613179905
	ctx := context.WithValue(req.Context(), ContextKeyID, id)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response response.FollowersResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, expected, response.Followers)
	assert.Equal(t, 200, w.Code)
}

func TestMyFollowings(t *testing.T) {
	expected := []response.Follower{
		{
			ID:    941150051613179905,
			Name:  "Chirag Makwana",
			Bio:   "Junior Software Engineer at ZURU TECH INDIA.",
			Image: nil,
		},
		{
			ID:    941445337509003265,
			Name:  "Sameer Makwana",
			Bio:   "Junior Software Engineer at ZURU TECH INDIA.",
			Image: nil,
		},
	}
	r.Get("/api/user/myfollwings", NewFollowerController(followerService).MyFollowings)

	req, _ := http.NewRequest("GET", "/api/user/myfollwings", http.NoBody)

	var id int64 = 942546435960995841
	ctx := context.WithValue(req.Context(), ContextKeyID, id)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response response.FollowingsResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, expected, response.Followings)
	assert.Equal(t, 200, w.Code)
}
