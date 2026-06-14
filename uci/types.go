package uci

// Dataset is a UCI Machine Learning Repository dataset.
type Dataset struct {
	Rank int    `json:"rank"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// wire types
type wireResponse struct {
	Status     int          `json:"status"`
	StatusText string       `json:"statusText"`
	Data       []wireDataset `json:"data"`
}

type wireDataset struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
