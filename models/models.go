package models

type PersonRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type PersonResponse struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int64  `json:"age"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
}

type AgeMarsh struct {
	Count int64  `json:"count"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
}

type GenderMarsh struct {
	Count       int64   `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type NationMarsh struct {
	Count   int64  `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
