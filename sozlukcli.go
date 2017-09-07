package sozlukcli

import (
	"fmt"
	"strconv"

	"github.com/gosimple/slug"
)

// Config :
type Config struct {
	token string
}

var config Config

// NewSozluk :
func NewSozluk(token string) bool {
	config = Config{token: token}

	response, err := isTokenExistReq(config.token)

	done := err == nil && response.Data.IsAlive

	if done {
		fmt.Printf("Hi %s!\n", response.Data.Slug)
	} else {
		fmt.Printf("This token (%s) not alive!\n - Please check config.json!\n", config.token)
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
