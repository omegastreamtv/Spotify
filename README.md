# Spotify

![Coverage](https://img.shields.io/badge/Coverage-81.5%25-brightgreen)

Spotify is a Go wrapper for Spotify's web API.

## Install

```cli
go get github.com/omegastreamtv/spotify
```

## Endpoints

### Albums

- [x] Get Album
- [x] Get Several Albums
- [x] Get Album Tracks
- [ ] Get User's Saved Albums
- [ ] Save Albums for Current User
- [ ] Remove Users' Saved Albums
- [ ] Check USer's Saved Albums
- [ ] Get New Releases

### Artists

- [x] Get Artist
- [x] Get Several Artists
- [x] Get Artist's Albums
- [x] Get Artist's Top Tracks
- [x] Get Artist's Related Artists

### Audiobooks

- [x] Get an Audiobook
- [x] Get Several Audiobooks
- [x] Get Audiobook Chapters
- [x] Get User's Saved Audiobooks
- [x] Save Audiobooks for Current User
- [x] Remove User's Saved Audiobooks
- [x] Check User's Saved Audiobooks

### Categories

- [x] Get Several Browse Categories
- [x] Get Single Browse Category

### Chapters

- [x] Get a Chapter
- [x] Get Several Chapters

### Episodes

- [x] Get Episode
- [x] Get Several Episodes
- [x] Get User's Saved Episodes
- [x] Save Episodes for Current User
- [x] Remove User's Saved Episodes
- [x] Check User's Saved Episodes

### Genres

- [x] Get Available Genre Seeds

### Markets

- [x] Get Available Markets

### Player

- [x] Get Playback State
- [ ] Transfer Playback
- [ ] Get Available Devices
- [ ] Get Currently Playing Track
- [ ] Start/Resume Playback
- [ ] Pause Playback
- [ ] Skip To Next
- [ ] Skip To Previous
- [ ] Skip to Position
- [ ] Set Repeat Mode
- [ ] Set Playback Mode
- [ ] Set Playback Volume
- [ ] Toggle Playback Shuffle
- [ ] Get Recently Played Tracks
- [ ] Get The User's Queue
- [ ] Add Item to Playback Queue

### Playlists

- [ ] Get Playlist
- [ ] Change Playlist Details
- [ ] Get Playlist Items
- [ ] Update Playlist Items
- [ ] Add Items to Playlist
- [ ] Remove Playlist Items
- [ ] Get Current User's Playlists
- [ ] Get User's Playlist
- [ ] Create Playlist
- [ ] Get Featured Playlists
- [ ] Get Category's Playlists
- [ ] Get Playlist Cover Image
- [ ] Add Custom Playlist Cover Image

### Search

- [x] Search for Item

### Shows

- [x] Get Show
- [x] Get Several Shows
- [x] Get Show Episodes
- [x] Get User's Saved Shows
- [x] Save Shows for Current User
- [x] Remove User's Saved Shows
- [x] Check User's Saved Shows

### Tracks

- [x] Get Track
- [x] Get Several Tracks
- [x] Get User's Saved Tracks
- [x] Save Tracks for Current User
- [x] Remove User's Saved Tracks
- [x] Check User's Saved Tracks
- [x] Get Tracks' Audio Features
- [x] Get Track's Audio Features
- [x] Get Track's Audio Analysis
- [x] Get Recommendations

### Users

- [x] Get Current User's Profile
- [x] Get User's Top Items (split into `GetUsersTopArtists` and `GetUsersTopTracks`)
- [x] Get User's Profile
- [x] Follow Playlist
- [x] Unfollow Playlist
- [x] Get Followed Artists
- [x] Follow Artists or Users
- [x] Unfollow Artists or Users
- [x] Check if User Follows Artists or Users
- [x] Check if Users Follow Playlist
