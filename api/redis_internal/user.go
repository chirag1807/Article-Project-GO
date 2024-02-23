package redis_internal

import (
	"articleproject/api/model/dto"
	"articleproject/api/model/request"
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func SaveUser(user request.User, id int64, rdb *redis.Client) {
	// redisUser := ConvertRequestuserToRedisUser(user, id)
	// redisUserMap := map[string]interface{}{
	// 	"id":   redisUser.ID,
	// 	"name": redisUser.Name,
	// 	"bio":   redisUser.Bio,
	// 	"email": redisUser.Email,
	// 	"image":   redisUser.Image,
	// 	"isadmin": redisUser.IsAdmin,
	// }
	key := "user:" + strconv.FormatInt(id, 10)

	// rdb.HSet(rdb.Context(), key, redisUserMap)
	rdb.HSet(rdb.Context(), key, "id", id, "name", user.Name, "bio", user.Bio, "email", user.Email, "image", user.Image, "isadmin", user.IsAdmin)
}

func ConvertRequestuserToRedisUser(user request.User, id int64) dto.User {
	return dto.User{
		ID:      id,
		Name:    user.Name,
		Bio:     user.Bio,
		Email:   user.Email,
		Image:   &user.Image,
		IsAdmin: user.IsAdmin,
	}
}

func GetUser(id int64, rdb *redis.Client) dto.User {
	key := "user:" + strconv.FormatInt(id, 10)
	user, _ := rdb.HGetAll(rdb.Context(), key).Result()
	var redisUser dto.User
	data, _ := json.Marshal(user)
	json.Unmarshal(data, &redisUser)
	return redisUser
}
