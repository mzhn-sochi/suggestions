package suggestions

type Suggestion struct {
	Title struct {
		Text string `json:"text"`
	} `json:"title"`
	Subtitle struct {
		Text string `json:"text"`
	} `json:"subtitle"`
	Tags     []string `json:"tags"`
	Distance struct {
		Value float64 `json:"value"`
		Text  string  `json:"text"`
	} `json:"distance"`
}
