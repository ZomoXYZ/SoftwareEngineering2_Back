package structs

type PlayerName struct {
	Adjective int `json:"adjective" binding:"required"`
	Noun int `json:"noun" binding:"required"`
}

type PlayerInfo struct {
	ID string `json:"id" binding:"required"`
	Name PlayerName `json:"name" binding:"required"`
	Picture int `json:"picture" binding:"required"`
}

type RestBodySelf struct {
	Name *PlayerName `json:"name,omitempty"`
	Picture *int `json:"picture,omitempty"`
}