package movies

import "github.com/google/uuid"

type MovieFilter struct {
	Id        *uuid.UUID `json:"id"`
	RatingGOE *int       `json:"ratingGOE"`
	TitleIn   []string   `json:"titleIn"`
	Search    string     `json:"search"`
}
