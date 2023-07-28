package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Either[L any, R any] struct {
	isLeft bool

	left  L
	right R
}

func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

type PlaylistTrack struct {
	AddedAt string                 `json:"added_at"`
	AddedBy User                   `json:"added_by"`
	IsLocal bool                   `json:"is_local"`
	Track   Either[Track, Episode] `json:"track"`
}

func EitherTrackOrEpisode[L Track, R Episode](data interface{}) (Either[Track, Episode], error) {
	e := Either[Track, Episode]{}

	// Try to convert the interface{} to map[string]interface{}
	m, ok := data.(map[string]interface{})
	if !ok {
		return e, errors.New("not a map")
	}

	// Convert map to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return e, err
	}

	// Try to unmarshal the JSON into a Track
	var tr Track
	err = json.Unmarshal(jsonData, &tr)
	if err == nil && tr.Name != "" { // Additional condition to ensure the struct is not empty
		e.isLeft = true
		e.left = tr
		return e, nil
	}

	// Try to unmarshal the JSON into an Episode
	var episode Episode
	err = json.Unmarshal(jsonData, &episode)
	if err == nil && episode.Name != "" { // Additional condition to ensure the struct is not empty
		fmt.Println("This is an Episode:", episode)
		e.isLeft = false
		e.right = episode
	}

	return e, nil
}
