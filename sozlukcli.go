package sozlukcli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)

// Config :
type Config struct {
	token string
	slug  string
}

var config Config

// NewSozluk :
func NewSozluk(token string) bool {
	config = Config{token: token, slug: ""}

	response, err := isTokenExistReq(config.token)

	done := err == nil && response.Data.IsAlive

	if done {
		config.slug = response.Data.Slug
	}

	return done
}

// CreateEntry :
func CreateEntry(topic string, entry string) int {
	topicID := TopicIsAlreadyExist(topic)

	if topicID == "" {
		response, err := createTopicWithEntryReq(topic, entry, config.token)

		done := err == nil && response.Success

		if done {
			return response.ID
		} else if err == nil && !response.Success {
			fmt.Printf("- %s\n", response.Message)
		}
	} else {
		response, err := createEntryIntoTopicReq(topicID, entry, config.token)

		done := err == nil && response.Success

		if done {
			return response.Data.ID
		} else if err == nil && !response.Success {
			fmt.Printf("- %s\n", response.Message)
		}
	}

	return -1
}

// DeleteEntry :
func DeleteEntry(ID int) {
	deleteEntryReq(ID, config.token)
}

// GetSlug :
func GetSlug() string {
	return config.slug
}

// TopicIsAlreadyExist :
func TopicIsAlreadyExist(topic string) string {
	response, err := isTopicExistReq(topic, config.token)

	done := err == nil && response.Success

	if done {
		target := slug.Make(topic)

		for index, element := range response.Data.Topics {
			_ = index
			if element.Slug == target {
				return strconv.Itoa((&element).ID)
			}
		}
	}

	return ""
}

// DoLogin :
func DoLogin(email string, password string) (string, error) {
	response, err := createSessionReq(email, password)

	done := err == nil && response.Success

	if done {
		return response.Data.Token, nil
	}

	return "", errors.New("Login Failed")
}

// DoLogout :
func DoLogout(token string) bool {
	deleteSessionReq(token)

	return true
}
