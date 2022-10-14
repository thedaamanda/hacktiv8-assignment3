package entity

type Status struct {
	Wind  int8 `json:"wind"`
	Water int8 `json:"water"`
}

type Data struct {
	Status Status `json:"status"`
}
