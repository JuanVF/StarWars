package controller

type Message struct {
	Name      string      `json:"username"`
	Text      string      `json:"text"`
	IdMessage string      `json:"idMessage"`
	Number    float64     `json:"number"`
	Numbers   []float64   `json:"numbers"`
	Matrix    [][]float64 `json:"matrix"`
}
