package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/omegastreamtv/spotify"
)

type App struct {
	client        *spotify.Client
	auth          *spotify.Authenticator
	state         string
	codeVerifier  string
	codeChallenge string
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
	app.codeVerifier = "ZTlrH7tD0nJJWe68uoTSqBWo5TA-jC4_LNW6sAm08EzqdgHK"
	app.codeChallenge = "xbqyIaXkWYQgYEp7n--EactN9WnvFbjSoKYBIhXgngs"
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
	http.Redirect(w, r, app.auth.AuthURL(
		app.state,
		spotify.ShowDialog,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", app.codeChallenge)),
		http.StatusFound,
	)
}

func (app *App) loginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.auth.Token(r.Context(), app.state, r,
		oauth2.SetAuthURLParam("code_verifier", app.codeVerifier))
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != app.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, app.state)
	}

	fmt.Println(token)
}
