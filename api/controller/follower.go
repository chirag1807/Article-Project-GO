package controller

import (
	"articleproject/api/model/response"
	"articleproject/api/service"
	"articleproject/constants"
	"articleproject/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type FollowerController interface {
	Follow(w http.ResponseWriter, r *http.Request)
	UnFollow(w http.ResponseWriter, r *http.Request)
	MyFollowers(w http.ResponseWriter, r *http.Request)
	MyFollowings(w http.ResponseWriter, r *http.Request)
	SomeoneFollowers(w http.ResponseWriter, r *http.Request)
	SomeoneFollowings(w http.ResponseWriter, r *http.Request)
}

type followerController struct {
	followerService service.FollowerService
}

func NewFollowerController(followerService service.FollowerService) FollowerController {
	return followerController{
		followerService: followerService,
	}
}

func (f followerController) Follow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := f.followerService.Follow(id, r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.FOLLOWING_NOW,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (f followerController) UnFollow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := f.followerService.UnFollow(id, r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.NOT_FOLLOWING_NOW,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (f followerController) MyFollowers(w http.ResponseWriter, r *http.Request) {
	followers, err := f.followerService.FetchFollowers(r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.FollowersResponse{
		Followers:    followers,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (f followerController) MyFollowings(w http.ResponseWriter, r *http.Request) {
	followings, err := f.followerService.FetchFollowings(r.Context().Value("id").(int64))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.FollowingsResponse{
		Followings: followings,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (f followerController) SomeoneFollowers(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	followers, err := f.followerService.FetchFollowers(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.FollowersResponse{
		Followers:    followers,
	}
	utils.ResponseGenerator(w, 200, response)
}

func (f followerController) SomeoneFollowings(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	followings, err := f.followerService.FetchFollowings(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.FollowingsResponse{
		Followings: followings,
	}
	utils.ResponseGenerator(w, 200, response)
}
