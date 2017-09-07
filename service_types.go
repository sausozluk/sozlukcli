package sozlukcli

// CheckResponse :
type CheckResponse struct {
	Success bool       `json:"success"`
	Data    *CheckData `json:"data"`
}

// CheckData :
type CheckData struct {
	IsAlive bool   `json: "isAlive"`
	User_Id string `json: "user_id"`
	Slug    string `json: "slug"`
	Unread  int    `json: "unread`
}

// SearchResponse :
type SearchResponse struct {
	Success bool        `json:"success"`
	Data    *SearchData `json:"data"`
}

// SearchData :
type SearchData struct {
	Topics []SearchDataTopics `json:"topics"`
	Users  []SearchDataUsers  `json:"users"`
}

// SearchDataTopics :
type SearchDataTopics struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// SearchDataUsers :
type SearchDataUsers struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
}

// EntryCreateResponse :
type EntryCreateResponse struct {
	Success bool             `json:"success"`
	Data    *EntryCreateData `json:"data"`
	Message string           `json:"message"`
}

// EntryCreateData :
type EntryCreateData struct {
	ID int `json:"id"`
}

// TopicCreateResponse :
type TopicCreateResponse struct {
	Success bool   `json:"success"`
	ID      int    `json:"entry_id"`
	Message string `json:"message"`
}
