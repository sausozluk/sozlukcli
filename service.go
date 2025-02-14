package sozlukcli

import (
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

const baseURL = "http://sausozluk.net/service/proxy/api/v1"

const (
	isTokenExistURL         = baseURL + "/sessions/check"
	isTopicExistURL         = baseURL + "/search?q="
	createTopicWithEntryURL = baseURL + "/topics"
	createEntryIntoTopicURL = baseURL + "/entries"
	deleteEntryURL          = baseURL + "/entries"
	createSessionURL        = baseURL + "/sessions"
	deleteSessionURL        = baseURL + "/sessions"
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

func deleteEntryReq(ID int, token string) (*EntryDeleteResponse, error) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetResult(EntryDeleteResponse{}).
		Delete(deleteEntryURL + "/" + strconv.Itoa(ID))

	response := resp.Result().(*EntryDeleteResponse)

	return response, err
}

func createSessionReq(email string, password string) (*SessionCreateResponse, error) {
	payload := fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		SetResult(SessionCreateResponse{}).
		Post(createSessionURL)

	response := resp.Result().(*SessionCreateResponse)

	return response, err
}

func deleteSessionReq(token string) (*SessionDeleteResponse, error) {
	resp, err := resty.R().
		SetHeader("token", token).
		SetHeader("Content-Type", "application/json").
		SetResult(SessionDeleteResponse{}).
		Delete(deleteSessionURL)

	response := resp.Result().(*SessionDeleteResponse)

	return response, err
}
