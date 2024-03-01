package main

import (
	testcase "articleproject/api/model/dto/test-cases"
	"articleproject/api/model/request"
	"articleproject/api/route"
	"articleproject/config"
	"articleproject/db"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

var conn *pgx.Conn
var rdb *redis.Client
var amqpConn *amqp.Connection

func runTestServer() *httptest.Server {
	config.LoadEnv("../.config/")
	var err error
	conn, rdb, amqpConn, err = db.DBConnection()
	if err != nil {
		log.Println(err)
		return nil
	}

	r := route.UsersRoutes(conn, rdb, amqpConn)

	return httptest.NewServer(r)
}

func Test_Post_Api_For_Add_Topic(t *testing.T) {
	buf := new(bytes.Buffer)
	var id int64 = 0
	ts := runTestServer()
	defer ts.Close()
	t.Run("it should return ok when topic insert done successfully.", func(t *testing.T) {
		topic := request.Topic{
			Name: "AWS",
		}
		jsonValue, _ := json.Marshal(topic)
		resp, _ := http.Post(fmt.Sprintf("%s/api/admin/add-topic", ts.URL), "application/json", bytes.NewBuffer(jsonValue))

		_, _ = buf.ReadFrom(resp.Body)
		responseBody := strings.TrimSuffix(buf.String(), "\n")
		var response testcase.IntegrationTestTopic
		json.Unmarshal([]byte(responseBody), &response)
		fmt.Println(responseBody)
		id = response.ID
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("it should return existed topics.", func(t *testing.T) {
		resp, _ := http.Get(fmt.Sprintf("%s/api/admin/get-all-topics", ts.URL))
		buf.Reset()
		buf.ReadFrom(resp.Body)
		responseBody := strings.TrimSuffix(buf.String(), "\n")
		fmt.Println(responseBody)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("it should return ok when topic update done successfully.", func(t *testing.T) {
		topic := request.Topic{
			ID:   id,
			Name: "Amazon Web Services",
		}
		jsonValue, _ := json.Marshal(topic)
		resp, _ := http.Post(fmt.Sprintf("%s/api/admin/update-topic", ts.URL), "application/json", bytes.NewBuffer(jsonValue))
		buf.Reset()
		_, _ = buf.ReadFrom(resp.Body)
		responseBody := strings.TrimSuffix(buf.String(), "\n")
		fmt.Println(responseBody)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("it should return ok when topic update done successfully.", func(t *testing.T) {
		resp, _ := http.Get(fmt.Sprintf("%s/api/admin/delete-topic/%d", ts.URL, id))
		buf.Reset()
		_, _ = buf.ReadFrom(resp.Body)
		responseBody := strings.TrimSuffix(buf.String(), "\n")
		fmt.Println(responseBody)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
