package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version,omitempty"`
}

type MovieGenres struct {
	Drama   string `json:"drama"`
	Romance string `json:"romance"`
	War     string `json:"war"`
}
