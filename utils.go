package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
)

func empty[T any]() (t T) {
	return
}

// Either respresents a value of 2 possible types.
// An instance of Either is an instance of either A or B.
type Either[L any, R any] struct {
	isLeft bool

	left  L
	right R
}

// IsLeft returns true if Either is an instance of Left.
func (e Either[L, R]) IsLeft() bool {
	return e.isLeft
}

// IsRight returns true if Either is an instance of Right.
func (e Either[L, R]) IsRight() bool {
	return !e.isLeft
}

// Left returns left value of a Either struct.
func (e Either[L, R]) Left() (L, bool) {
	if e.IsLeft() {
		return e.left, true
	}
	return empty[L](), false
}

// Right returns right value of a Either struct.
func (e Either[L, R]) Right() (R, bool) {
	if e.IsRight() {
		return e.right, true
	}
	return empty[R](), false
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

	return e, errors.New("data could not be parsed as a Track or Episode")
}
