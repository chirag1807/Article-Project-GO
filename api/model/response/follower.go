package response

type Follower struct {
	ID       int64      `json:"id,omitempty"`
	Name     string     `json:"name"`
	Bio      string     `json:"bio"`
	Image    *string 	`json:"image"`
}

type FollowersResponse struct {
	Followers []Follower
}

type FollowingsResponse struct {
	Followings []Follower
}