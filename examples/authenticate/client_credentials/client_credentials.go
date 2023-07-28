package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/omegastreamtv/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	ctx := context.Background()

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client, err := spotify.NewClient(
		spotify.WithClientID(os.Getenv("SPOTIFY_ID")),
		spotify.WithClientSecret(os.Getenv("SPOTIFY_SECRET")),
		spotify.WithRedirectURI("http://localhost:8080/callback"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client.SetAppAccessToken(token.AccessToken)

	singleAlbum, err := client.GetAlbum("382ObEPsp2rxGrnsizN5TX", "es")
	if err != nil {
		if spotifyErr, ok := err.(*spotify.SpotifyError); ok {
			fmt.Println(spotifyErr.Err.Status, spotifyErr.Err.Message)
		}

		log.Fatal(err)
		return
	}

	fmt.Println(singleAlbum)
}
