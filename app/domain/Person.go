package domain

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}
