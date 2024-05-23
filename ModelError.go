package credmark

type ModelError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Stack   []struct {
		Slug string `json:"slug"`
	} `json:"stack"`
}
