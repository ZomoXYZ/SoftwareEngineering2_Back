package structs

type PlayerName struct {
	Adjective int `json:"adjective"`
	Noun int `json:"noun"`
}

type Player struct {
	ID string `json:"id"`
	Name PlayerName `json:"name"`
	Picture int `json:"picture"`
}

type PlayerNameOpt struct {
	Adjective *int `json:"adjective" validate:"exists"`
	Noun *int `json:"noun" validate:"exists"`
}

type RestBodySelf struct {
	Name *PlayerNameOpt `json:"name"`
	Picture *int `json:"picture"`
}