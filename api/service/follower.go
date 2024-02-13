package service

import (
	"articleproject/api/model/response"
	"articleproject/api/repository"
)

type FollowerService interface {
	Follow(int64, int64) error
	UnFollow(int64, int64) error
	FetchFollowers(int64) ([]response.Follower, error)
	FetchFollowings(int64) ([]response.Follower, error)
}

type followerService struct {
	followerRepository repository.FollowerRepository
}

func NewFollowerService(followerRepository repository.FollowerRepository) FollowerService {
	return followerService{
		followerRepository: followerRepository,
	}
}

func (f followerService) Follow(userid int64, followersid int64) error {
	return f.followerRepository.Follow(userid, followersid)
}

func (f followerService) UnFollow(userid int64, followersid int64) error {
	return f.followerRepository.UnFollow(userid, followersid)
}

func (f followerService) FetchFollowers(userid int64) ([]response.Follower, error) {
	return f.followerRepository.FetchFollowers(userid)
}

func (f followerService) FetchFollowings(followerid int64) ([]response.Follower, error) {
	return f.followerRepository.FetchFollowings(followerid)
}