package repository

import (
	"articleproject/api/model/request"
	"articleproject/api/model/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTopic(t *testing.T) {
	topic := request.Topic{
		Name: "React-Native",
	}
	err := NewTopicRepo(conn).AddTopic(topic)
	assert.Equal(t, nil, err)
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
			ID:   947139268841570305,
			Name: "React-Native",
		},
	}
	topics, _ := NewTopicRepo(conn).GetAllTopics()

	assert.Equal(t, expected, topics)
}

func TestUpdateTopic(t *testing.T) {
	topic := request.Topic{
		ID:   947139268841570305,
		Name: "ReactNative",
	}

	err := NewTopicRepo(conn).UpdateTopic(topic)
	assert.Equal(t, nil, err)
}

func TestDeleteTopic(t *testing.T) {
	err := NewTopicRepo(conn).DeleteTopic(947139268841570305)
	assert.Equal(t, nil, err)
}
