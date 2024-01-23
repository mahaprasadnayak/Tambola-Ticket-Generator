package utils

type Ticket struct {
	Numbers [][]int `json:"numbers"`
}

type TambolaSet struct {
	Tickets []Ticket `json:"tickets"`
}
