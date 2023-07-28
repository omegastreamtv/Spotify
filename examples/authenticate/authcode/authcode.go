package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/omegastreamtv/spotify"
)

type App struct {
	client *spotify.Client
	auth   *spotify.Authenticator
	state  string
}

func main() {
	client, err := spotify.NewClient(
		spotify.WithClientID(os.Getenv("SPOTIFY_ID")),
		spotify.WithClientSecret(os.Getenv("SPOTIFY_SECRET")),
		spotify.WithRedirectURI("http://localhost:8080/callback"),
	)
	if err != nil {
		panic(err)
	}

	app := &App{client: client}
	app.state = "some-random-state-generated-each-request"
	app.auth = client.NewAuthenticator(
		spotify.WithScopes(
			spotify.ScopeUserReadEmail,
			spotify.ScopeUserReadPlaybackState,
			spotify.ScopeUserReadRecentlyPlayed,
		),
	)

	http.HandleFunc("/login", app.loginHandler)
	http.HandleFunc("/callback", app.loginCallbackHandler)
	http.ListenAndServe(":8080", nil)
}

func (app *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, app.auth.AuthURL(app.state, spotify.ShowDialog), http.StatusFound)
}

func (app *App) loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.auth.Token(r.Context(), app.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		panic(err)
	}

	if st := r.FormValue("state"); st != app.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, app.state)
	}

	fmt.Println(token)
}
