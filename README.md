# Spotify

![Coverage](https://img.shields.io/badge/Coverage-75.5%25-brightgreen)

Spotify is a Go wrapper for Spotify's web API.

## Install

```cli
go get github.com/omegastreamtv/spotify
```

## Endpoints

_All the `Params`, `Body` and `Response` types are named exactly the same as the function._

### Albums

- [x] [Get Album (`GetAlbum`)](https://developer.spotify.com/documentation/web-api/reference/get-an-album)
- [x] [Get Several Albums (`GetSeveralAlbums`)](https://developer.spotify.com/documentation/web-api/reference/get-multiple-albums)
- [x] [Get Album Tracks (`GetAlbumTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-an-albums-tracks)
- [x] [Get User's Saved Albums (`GetUsersSavedAlbums`)](https://developer.spotify.com/documentation/web-api/reference/get-users-saved-albums)
- [x] [Save Albums for Current User (`SaveAlbumsForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/save-albums-user)
- [x] [Remove Users' Saved Albums (`RemoveAlbumsForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/remove-albums-user)
- [x] [Check USer's Saved Albums (`CheckUsersSavedAlbums`)](https://developer.spotify.com/documentation/web-api/reference/check-users-saved-albums)
- [x] [Get New Releases (`GetNewReleases`)](https://developer.spotify.com/documentation/web-api/reference/get-new-releases)

### Artists

- [x] [Get Artist (`GetArtist`)](https://developer.spotify.com/documentation/web-api/reference/get-an-artist)
- [x] [Get Several Artists (`GetSeveralArtists`)](https://developer.spotify.com/documentation/web-api/reference/get-multiple-artists)
- [x] [Get Artist's Albums (`GetArtistsAlbums`)](https://developer.spotify.com/documentation/web-api/reference/get-an-artists-albums)
- [x] [Get Artist's Top Tracks (`GetArtistsTopTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-an-artists-top-tracks)
- [x] [Get Artist's Related Artists (`GetArtistsRelatedArtists`)](https://developer.spotify.com/documentation/web-api/reference/get-an-artists-related-artists)

### Audiobooks

- [x] [Get an Audiobook (`GetAnAudiobook`)](https://developer.spotify.com/documentation/web-api/reference/get-an-audiobook)
- [x] [Get Several Audiobooks (`GetSeveralAudiobooks`)](https://developer.spotify.com/documentation/web-api/reference/get-multiple-audiobooks)
- [x] [Get Audiobook Chapters (`GetAudiobookChapters`)](https://developer.spotify.com/documentation/web-api/reference/get-audiobook-chapters)
- [x] [Get User's Saved Audiobooks (`GetUsersSavedAudioBooks`)](https://developer.spotify.com/documentation/web-api/reference/get-users-saved-audiobooks)
- [x] [Save Audiobooks for Current User (`SaveAudiobooksForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/save-audiobooks-user)
- [x] [Remove User's Saved Audiobooks (`RemoveUsersSavedAudiobooks`)](https://developer.spotify.com/documentation/web-api/reference/remove-audiobooks-user)
- [x] [Check User's Saved Audiobooks (`CheckUsersSavedAudiobooks`)](https://developer.spotify.com/documentation/web-api/reference/check-users-saved-audiobooks)

### Categories

- [x] [Get Several Browse Categories (`GetSeveralBrowseCategories`)](https://developer.spotify.com/documentation/web-api/reference/get-categories)
- [x] [Get Single Browse Category (`GetSingleBrowseCategory`)](https://developer.spotify.com/documentation/web-api/reference/get-a-category)

### Chapters

- [x] [Get a Chapter (`GetAChapter`)](https://developer.spotify.com/documentation/web-api/reference/get-a-chapter)
- [x] [Get Several Chapters (`GetSeveralChapters`)](https://developer.spotify.com/documentation/web-api/reference/get-several-chapters)

### Episodes

- [x] [Get Episode (`GetEpisode`)](https://developer.spotify.com/documentation/web-api/reference/get-an-episode)
- [x] [Get Several Episodes (`GetSeveralEpisodes`)](<[url](https://developer.spotify.com/documentation/web-api/reference/get-multiple-episodes)>)
- [x] [Get User's Saved Episodes (`GetUsersSavedEpisodes`)](https://developer.spotify.com/documentation/web-api/reference/get-users-saved-episodes)
- [x] [Save Episodes for Current User (`SaveEpisodesForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/save-episodes-user)
- [x] [Remove User's Saved Episodes (`RemoveUsersSavedEpisodes`)](https://developer.spotify.com/documentation/web-api/reference/remove-episodes-user)
- [x] [Check User's Saved Episodes (`CheckUsersSavedEpisodes`)](https://developer.spotify.com/documentation/web-api/reference/check-users-saved-episodes)

### Genres

- [x] [Get Available Genre Seeds (`GetAvailableGenreSeeds`)](https://developer.spotify.com/documentation/web-api/reference/get-recommendation-genres)

### Markets

- [x] [Get Available Markets (`GetAvailableMarkets`)](https://developer.spotify.com/documentation/web-api/reference/get-available-markets)

### Player

- [x] [Get Playback State (`GetPlaybackState`)](https://developer.spotify.com/documentation/web-api/reference/get-information-about-the-users-current-playback)
- [ ] [Transfer Playback (`TransferPlayback`)](https://developer.spotify.com/documentation/web-api/reference/transfer-a-users-playback)
- [ ] [Get Available Devices (`GetAvailableDevices`)](https://developer.spotify.com/documentation/web-api/reference/get-a-users-available-devices)
- [ ] [Get Currently Playing Track (`GetCurrentlyPlayingTrack`)](https://developer.spotify.com/documentation/web-api/reference/get-the-users-currently-playing-track)
- [ ] [Start/Resume Playback (`StartResumePlayback`)](https://developer.spotify.com/documentation/web-api/reference/start-a-users-playback)
- [ ] [Pause Playback (`PausePlayback`)](https://developer.spotify.com/documentation/web-api/reference/pause-a-users-playback)
- [ ] [Skip To Next (`SkipToNext`)](https://developer.spotify.com/documentation/web-api/reference/skip-users-playback-to-next-track)
- [ ] [Skip To Previous (`SkipToPrevious`)](https://developer.spotify.com/documentation/web-api/reference/skip-users-playback-to-previous-track)
- [ ] [Skip to Position (`SeekToPosition`)](https://developer.spotify.com/documentation/web-api/reference/seek-to-position-in-currently-playing-track)
- [ ] [Set Repeat Mode (`SetRepeatMode`)](https://developer.spotify.com/documentation/web-api/reference/set-repeat-mode-on-users-playback)
- [ ] [Set Playback Volume (`SetPlaybackVolume`)](https://developer.spotify.com/documentation/web-api/reference/set-volume-for-users-playback)
- [ ] [Toggle Playback Shuffle (`TogglePlaybackShuffle`)](https://developer.spotify.com/documentation/web-api/reference/toggle-shuffle-for-users-playback)
- [ ] [Get Recently Played Tracks (`GetRecentlyPlayedTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-recently-played)
- [ ] [Get The User's Queue (`GetTheUsersQueue`)](https://developer.spotify.com/documentation/web-api/reference/get-queue)
- [ ] [Add Item to Playback Queue (`AddItemToPlaybackQueue`)](https://developer.spotify.com/documentation/web-api/reference/add-to-queue)

### Playlists

- [ ] [Get Playlist (`GetPlaylist`)](https://developer.spotify.com/documentation/web-api/reference/get-playlist)
- [ ] [Change Playlist Details (`ChangePlaylistDetails`)](https://developer.spotify.com/documentation/web-api/reference/change-playlist-details)
- [ ] [Get Playlist Items (`GetPlaylistItems`)](https://developer.spotify.com/documentation/web-api/reference/get-playlists-tracks)
- [ ] [Update Playlist Items (`UpdatePlaylistItems`)](https://developer.spotify.com/documentation/web-api/reference/reorder-or-replace-playlists-tracks)
- [ ] [Add Items to Playlist (`AddItemsToPlaylist`)](https://developer.spotify.com/documentation/web-api/reference/add-tracks-to-playlist)
- [ ] [Remove Playlist Items (`RemovePlaylistItems`)](https://developer.spotify.com/documentation/web-api/reference/remove-tracks-playlist)
- [ ] [Get Current User's Playlists (`GetCurrentUsersPlaylists`)](https://developer.spotify.com/documentation/web-api/reference/get-a-list-of-current-users-playlists)
- [ ] [Get User's Playlist (`GetUsersPlaylists`)](https://developer.spotify.com/documentation/web-api/reference/get-list-users-playlists)
- [ ] [Create Playlist (`CreatePlaylist`)](https://developer.spotify.com/documentation/web-api/reference/create-playlist)
- [ ] [Get Featured Playlists (`GetFeaturedPlaylists`)](https://developer.spotify.com/documentation/web-api/reference/get-featured-playlists)
- [ ] [Get Category's Playlists (`GetCategorysPlaylists`)](https://developer.spotify.com/documentation/web-api/reference/get-a-categories-playlists)
- [ ] [Get Playlist Cover Image (`GetPlaylistCoverImage`)](https://developer.spotify.com/documentation/web-api/reference/get-playlist-cover)
- [ ] [Add Custom Playlist Cover Image (`AddCustomPlaylistCoverImage`)](https://developer.spotify.com/documentation/web-api/reference/upload-custom-playlist-cover)

### Search

- [x] [Search for Item (`SearchForItem`)](https://developer.spotify.com/documentation/web-api/reference/search)

### Shows

- [x] [Get Show (`GetShow`)](https://developer.spotify.com/documentation/web-api/reference/get-a-show)
- [x] [Get Several Shows (`GetSeveralShows`)](https://developer.spotify.com/documentation/web-api/reference/get-multiple-shows)
- [x] [Get Show Episodes (`GetShowEpisodes`)](https://developer.spotify.com/documentation/web-api/reference/get-a-shows-episodes)
- [x] [Get User's Saved Shows (`GetUsersSavedShows`)](https://developer.spotify.com/documentation/web-api/reference/get-users-saved-shows)
- [x] [Save Shows for Current User (`SaveShowsForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/save-shows-user)
- [x] [Remove User's Saved Shows (`RemoveUsersSavedShows`)](https://developer.spotify.com/documentation/web-api/reference/remove-shows-user)
- [x] [Check User's Saved Shows (`CheckUsersSavedShows`)](https://developer.spotify.com/documentation/web-api/reference/check-users-saved-shows)

### Tracks

- [x] [Get Track (`GetTrack`)](https://developer.spotify.com/documentation/web-api/reference/get-track)
- [x] [Get Several Tracks (`GetSeveralTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-several-tracks)
- [x] [Get User's Saved Tracks (`GetUsersSavedTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-users-saved-tracks)
- [x] [Save Tracks for Current User (`SaveTracksForCurrentUser`)](https://developer.spotify.com/documentation/web-api/reference/save-tracks-user)
- [x] [Remove User's Saved Tracks (`RemoveUsersSavedTracks`)](https://developer.spotify.com/documentation/web-api/reference/remove-tracks-user)
- [x] [Check User's Saved Tracks (`CheckUsersSavedTracks`)](https://developer.spotify.com/documentation/web-api/reference/check-users-saved-tracks)
- [x] [Get Tracks' Audio Features (`GetMultiTracksAudioFeatures`)](https://developer.spotify.com/documentation/web-api/reference/get-several-audio-features)
- [x] [Get Track's Audio Features (`GetSingleTracksAudioFeatures`)](https://developer.spotify.com/documentation/web-api/reference/get-audio-features)
- [x] [Get Track's Audio Analysis (`GetTracksAudioAnalysis`)](https://developer.spotify.com/documentation/web-api/reference/get-audio-analysis)
- [x] [Get Recommendations (`GetRecommendations`)](https://developer.spotify.com/documentation/web-api/reference/get-recommendations)

### Users

- [x] [Get Current User's Profile (`GetCurrentUsersProfile`)](https://developer.spotify.com/documentation/web-api/reference/get-current-users-profile)
- [x] [Get User's Top Items (split into `GetUsersTopArtists` and `GetUsersTopTracks`)](https://developer.spotify.com/documentation/web-api/reference/get-users-top-artists-and-tracks)
- [x] [Get User's Profile (`GetUsersProfile`)](https://developer.spotify.com/documentation/web-api/reference/get-users-profile)
- [x] [Follow Playlist (`FollowPlaylist`)](https://developer.spotify.com/documentation/web-api/reference/follow-playlist)
- [x] [Unfollow Playlist (`UnfollowPlaylist`)](https://developer.spotify.com/documentation/web-api/reference/unfollow-playlist)
- [x] [Get Followed Artists (`GetFollowedArtists`)](https://developer.spotify.com/documentation/web-api/reference/get-followed)
- [x] [Follow Artists or Users (`FollowArtistsOrUsers`)](https://developer.spotify.com/documentation/web-api/reference/follow-artists-users)
- [x] [Unfollow Artists or Users (`UnfollowArtistsOrUsers`)](https://developer.spotify.com/documentation/web-api/reference/unfollow-artists-users)
- [x] [Check if User Follows Artists or Users (`CheckIfUserFollowsArtistsOrUsers`)](https://developer.spotify.com/documentation/web-api/reference/check-current-user-follows)
- [x] [Check if Users Follow Playlist (`CheckIfUsersFollowPlaylist`)](https://developer.spotify.com/documentation/web-api/reference/check-if-user-follows-playlist)
