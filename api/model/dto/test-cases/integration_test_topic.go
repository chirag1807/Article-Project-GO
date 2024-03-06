package testcase

type IntegrationTestTopic struct {
	ID    int64  `json:"id"`
	Error *error `json:"err,omitempty"`

}