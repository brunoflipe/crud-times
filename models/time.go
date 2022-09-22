package models

var ListaDeTimes []*Time

type Time struct {
	ID      uint   `json:"id"`
	Nome    string `json:"nome"`
	Titulos uint   `json:"titulos"`
	Tecnico string `json:"tecnico"`
}
