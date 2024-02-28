package controller

import (
	testcase "articleproject/api/model/dto/test-cases"
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestAddTopic(t *testing.T) {
	r.Post("/api/admin/add-topic", NewTopicController(topicService).AddTopic)
	topic := request.Topic{
		Name: "React-Native",
	}
	jsonValue, _ := json.Marshal(topic)

	req, _ := http.NewRequest("POST", "/api/admin/add-topic", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var response testcase.Topic
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Topic Added Successfully.", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestGetAllTopics(t *testing.T) {
	expected := []response.Topic{
		{
			ID:   941510542024474625,
			Name: "GO",
		},
		{
			ID:   941510596722393089,
			Name: "Node.JS",
		},
		{
			ID:   941510622923816961,
			Name: "Flutter",
		},
		{
			ID:   943765517224148993,
			Name: "React",
		},
		{
			ID:   947088469108785153,
			Name: "React-Native",
		},
	}
	r.Get("/api/admin/get-all-topics", NewTopicController(topicService).GetAllTopics)
	req, _ := http.NewRequest("GET", "/api/admin/get-all-topics", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var response response.TopicResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, expected, response.Topics)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateTopic(t *testing.T) {
	r.Put("/api/admin/update-topic", NewTopicController(topicService).UpdateTopic)

	topic := request.Topic{
		ID:   947088469108785153,
		Name: "ReactNative",
	}
	jsonValue, _ := json.Marshal(topic)
	req, _ := http.NewRequest("PUT", "/api/admin/update-topic", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	var response testcase.Topic
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Topic Updated Successfully.", response.Message)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteTopic(t *testing.T) {
	r.Delete("/api/admin/delete-topic/:ID", NewTopicController(topicService).DeleteTopic)

	req, _ := http.NewRequest("DELETE", "/api/admin/delete-topic/:ID", http.NoBody)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("ID", "947088469108785153")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	r.ServeHTTP(w, req)

	var response testcase.Topic
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Topic Deleted Successfully.", response.Message)
	assert.Equal(t, 200, w.Code)
}
