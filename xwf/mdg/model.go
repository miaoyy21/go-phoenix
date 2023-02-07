package mdg

type Model struct {
	Nodes []struct {
		Text string `json:"text"`
		Key  int    `json:"key"`
	} `json:"nodeDataArray"`

	Links []struct {
		From int `json:"from"`
		To   int `json:"to"`
	} `json:"linkDataArray"`
}
