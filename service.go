package sozlukcli

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

const (
	isTokenExistURL         = "http://sausozluk.net/service/proxy/api/v1/sessions/check"
	isTopicExistURL         = "http://sausozluk.net/service/proxy/api/v1/search?q="
	createTopicWithEntryURL = "http://sausozluk.net/service/proxy/api/v1/topics"
	createEntryIntoTopicURL = "http://sausozluk.net/service/proxy/api/v1/entries"
)

func isTokenExistReq(token string) (*CheckResponse, error) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetResult(CheckResponse{}).
		Post(isTokenExistURL)

	response := resp.Result().(*CheckResponse)

	return response, err
}

func isTopicExistReq(topic string, token string) (*SearchResponse, error) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetResult(SearchResponse{}).
		Get(isTopicExistURL + topic)

	response := resp.Result().(*SearchResponse)

	return response, err
}

func createTopicWithEntryReq(topic string, entry string, token string) (*TopicCreateResponse, error) {
	payload := fmt.Sprintf(`{"topic": {"title": "%s"}, "entry": {"text": "%s"}}`, topic, entry)

	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(TopicCreateResponse{}).
		Post(createTopicWithEntryURL)

	response := resp.Result().(*TopicCreateResponse)

	return response, err
}

func createEntryIntoTopicReq(topicID string, entry string, token string) (*EntryCreateResponse, error) {
	payload := fmt.Sprintf(`{"topic_id": "%s", "text": "%s"}`, topicID, entry)

	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(EntryCreateResponse{}).
		Post(createEntryIntoTopicURL)

	response := resp.Result().(*EntryCreateResponse)

	return response, err
}
