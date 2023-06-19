package models

type RequestPersianQa struct {
	Text      string            `json:"text"`
	Questions map[string]string `json:"questions"`
}

type OutputPersianQa struct {
	Text string `json:"text"`
	QA   []struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"qa"`
}
