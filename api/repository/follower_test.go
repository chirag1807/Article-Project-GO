package repository

import (
	"articleproject/api/model/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	err := NewFollowerRepository(conn).Follow(941445337509003265, 941150051613179905)
	assert.Equal(t, nil, err)
}

func TestUnFollow(t *testing.T) {
	err := NewFollowerRepository(conn).UnFollow(941445337509003265, 941150051613179905)
	assert.Equal(t, nil, err)
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
	
	followers, _ := NewFollowerRepository(conn).FetchFollowers(941150051613179905)
	assert.Equal(t, expected, followers)
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

	followings, _ := NewFollowerRepository(conn).FetchFollowings(942546435960995841)
	assert.Equal(t, expected, followings)
}
