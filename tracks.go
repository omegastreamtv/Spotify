package spotify

import (
	"fmt"
	"strings"
)

type Track struct {
	// The album on which the track appears. The album object includes a link in href to full information about the album.
	Album Album `json:"album"`
	// The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	Artists []Artist `json:"artists"`
	// A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually 1 unless the album consists of more than one disc).
	DiscNumber int `json:"disc_number"`
	// The track length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// Known external IDs for the track.
	ExternalIDs ExternalIDs `json:"external_ids"`
	// Known external URLs for this track.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The Spotify ID for the track.
	ID string `json:"id"`
	// Part of the response when Track Relinking is applied. If true, the track is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`
	// Part of the response when Track Relinking is applied, and the requested track has been replaced with different track.
	//
	// The track in the linked_from object contains information about the originally requested track.
	LinkedFrom LinkedFrom `json:"linked_from"`
	// Included in the response when a content restriction is applied.
	Restrictions Restrictions `json:"restrictions"`
	// The name of the track.
	Name string `json:"name"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	//
	// The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.
	//
	// Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. Note: the popularity value may lag actual popularity by a few days: the value is not updated in real time.
	Popularity int `json:"popularity"`
	// A link to a 30 second preview (MP3 format) of the track. Can be null
	PreviewURL string `json:"preview_url,omitempty"`
	// The number of the track. If an album has several discs, the track number is the number on the specified disc.
	TrackNumber int `json:"track_number"`
	// The object type: "track".
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
}

type LinkedFrom struct {
	// Known external URLs for this track.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The Spotify ID for the track.
	ID string `json:"id"`
	// The object type: "track".
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
}

type SavedTrack struct {
	// The date and time the track was saved. Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.
	AddedAt string `json:"added_at"`
	Track   Track  `json:"track"`
}

type AudioFeatures struct {
	// A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic.
	Acousticness float32 `json:"acousticness"`
	// A URL to access the full audio analysis of this track. An access token is required to access this data.
	AnalysisURL string `json:"analysis_url"`
	// Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable.
	Danceability float32 `json:"danceability"`
	// The duration of the track in milliseconds.
	DurationMS int `json:"duration_ms"`
	// Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. For example, death metal has high energy, while a Bach prelude scores low on the scale. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy.
	Energy float32 `json:"energy"`
	// The Spotify ID for the track.
	ID string `json:"id"`
	// Predicts whether a track contains no vocals. "Ooh" and "aah" sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly "vocal". The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0.
	Instrumentalness float32 `json:"instrumentalness"`
	// The key the track is in. Integers map to pitches using standard Pitch Class notation. E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.
	Key int `json:"key"`
	// Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live.
	Liveness float32 `json:"liveness"`
	// The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.
	Loudness float32 `json:"loudness"`
	// Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.
	Mode int `json:"mode"`
	// Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks.
	Speechiness float32 `json:"speachiness"`
	// The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.
	Tempo float32 `json:"tempo"`
	// An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
	TimeSignature int `json:"time_signature"`
	// A link to the Web API endpoint providing full details of the track.
	TrackHref string `json:"track_href"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
	// A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry).
	Valence float32 `json:"valence"`
}

type AudioAnalysis struct {
	Meta struct {
		// The version of the Analyzer used to analyze this track.
		AnalyzerVersion string `json:"analyzer_version"`
		// The platform used to read the track's audio data.
		Platform string `json:"platform"`
		// A detailed status code for this track. If analysis data is missing, this code may explain why.
		DetailedStatus string `json:"detailed_status"`
		// The return code of the analyzer process. 0 if successful, 1 if any errors occurred.
		StatusCode int `json:"status_code"`
		// The Unix timestamp (in seconds) at which this track was analyzed.
		Timestamp int `json:"timestamp"`
		// The amount of time taken to analyze this track.
		AnalysisTime float32 `json:"analysis_time"`
		// The method used to read the track's audio data.
		InputProcess string `json:"input_process"`
	} `json:"meta"`
	Track struct {
		// The exact number of audio samples analyzed from this track. See also analysis_sample_rate.
		NumSamples int `json:"num_samples"`
		// Length of the track in seconds.
		Duration float32 `json:"duration"`
		// This field will always contain the empty string.
		SampleMD5 string `json:"sample_md5"`
		// An offset to the start of the region of the track that was analyzed. (As the entire track is analyzed, this should always be 0.)
		OffsetSeconds int `json:"offset_seconds"`
		// The length of the region of the track was analyzed, if a subset of the track was analyzed. (As the entire track is analyzed, this should always be 0.)
		WindowSeconds int `json:"window_seconds"`
		// The sample rate used to decode and analyze this track. May differ from the actual sample rate of this track available on Spotify.
		AnalysisSampleRate int `json:"analysis_sample_rate"`
		// The number of channels used for analysis. If 1, all channels are summed together to mono before analysis.
		AnalysisChannels int `json:"analysis_channels"`
		// The time, in seconds, at which the track's fade-in period ends. If the track has no fade-in, this will be 0.0.
		EndOfFadeIn float32 `json:"end_of_fade_in"`
		// The time, in seconds, at which the track's fade-out period starts. If the track has no fade-out, this should match the track's length.
		StartOfFadeOut float32 `json:"start_of_fade_out"`
		// The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.
		Loudness float32 `json:"loudness"`
		// The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.
		Tempo float32 `json:"tempo"`
		// The confidence, from 0.0 to 1.0, of the reliability of the tempo.
		TempoConfidence float32 `json:"tempo_confidence"`
		// An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
		TimeSignature int `json:"time_signature"`
		// The confidence, from 0.0 to 1.0, of the reliability of the time_signature.
		TimeSignatureConfidence float32 `json:"time_signature_confidence"`
		// The key the track is in. Integers map to pitches using standard Pitch Class notation. E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1.
		Key int `json:"key"`
		// The confidence, from 0.0 to 1.0, of the reliability of the key.
		KeyConfidence float32 `json:"key_confidence"`
		// Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.
		Mode int `json:"mode"`
		// The confidence, from 0.0 to 1.0, of the reliability of the mode.
		ModeConfidence float32 `json:"mode_confidence"`
		// An Echo Nest Musical Fingerprint (ENMFP) codestring for this track.
		Codestring string `json:"codestring"`
		// A version number for the Echo Nest Musical Fingerprint format used in the codestring field.
		CodeVersion float32 `json:"code_version"`
		// An EchoPrint codestring for this track.
		Echoprintstring string `json:"echoprintstring"`
		// A version number for the EchoPrint format used in the echoprintstring field.
		EchoprintVersion float32 `json:"echoprint_version"`
		// A Synchstring for this track.
		Synchstring string `json:"synchstring"`
		// A version number for the Synchstring used in the synchstring field.
		SynchVersion float32 `json:"synch_version"`
		// A Rhythmstring for this track. The format of this string is similar to the Synchstring.
		Rhythmstring string `json:"rhythmstring"`
		// A version number for the Rhythmstring used in the rhythmstring field.
		RhythmVersion float32 `json:"rhythm_version"`
	} `json:"track"`
	// The time intervals of the bars throughout the track. A bar (or measure) is a segment of time defined as a given number of beats.
	Bars []struct {
		// The starting point (in seconds) of the time interval.
		Start float32 `json:"start"`
		// The duration (in seconds) of the time interval.
		Duration float32 `json:"duration"`
		// The confidence, from 0.0 to 1.0, of the reliability of the interval.
		Confidence float32 `json:"confidence"`
	} `json:"bars"`
	// The time intervals of beats throughout the track. A beat is the basic time unit of a piece of music; for example, each tick of a metronome. Beats are typically multiples of tatums.
	Beats []struct {
		// The starting point (in seconds) of the time interval.
		Start float32 `json:"start"`
		// The duration (in seconds) of the time interval.
		Duration float32 `json:"duration"`
		// The confidence, from 0.0 to 1.0, of the reliability of the interval.
		Confidence float32 `json:"confidence"`
	} `json:"beats"`
	// Sections are defined by large variations in rhythm or timbre, e.g. chorus, verse, bridge, guitar solo, etc. Each section contains its own descriptions of tempo, key, mode, time_signature, and loudness.
	Sections []struct {
		// The starting point (in seconds) of the time interval.
		Start float32 `json:"start"`
		// The duration (in seconds) of the time interval.
		Duration float32 `json:"duration"`
		// The confidence, from 0.0 to 1.0, of the reliability of the interval.
		Confidence float32 `json:"confidence"`
		// The overall loudness of the section in decibels (dB). Loudness values are useful for comparing relative loudness of sections within tracks.
		Loudness float32 `json:"loudness"`
		// The overall estimated tempo of the section in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.
		Tempo float32 `json:"tempo"`
		// The confidence, from 0.0 to 1.0, of the reliability of the tempo. Some tracks contain tempo changes or sounds which don't contain tempo (like pure speech) which would correspond to a low value in this field.
		TempoConfidence float32 `json:"tempo_confidence"`
		// The estimated overall key of the section. The values in this field ranging from 0 to 11 mapping to pitches using standard Pitch Class notation (E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on). If no key was detected, the value is -1.
		Key int `json:"key"`
		// The confidence, from 0.0 to 1.0, of the reliability of the key. Songs with many key changes may correspond to low values in this field.
		KeyConfidence float32 `json:"key_confidence"`
		// Indicates the modality (major or minor) of a section, the type of scale from which its melodic content is derived. This field will contain a 0 for "minor", a 1 for "major", or a -1 for no result. Note that the major key (e.g. C major) could more likely be confused with the minor key at 3 semitones lower (e.g. A minor) as both keys carry the same pitches.
		Mode int `json:"mode"`
		// The confidence, from 0.0 to 1.0, of the reliability of the mode.
		ModeConfidence float32 `json:"mode_confidence"`
		// An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of "3/4", to "7/4".
		TimeSignature int `json:"time_signature"`
		// The confidence, from 0.0 to 1.0, of the reliability of the time_signature. Sections with time signature changes may correspond to low values in this field.
		TimeSignatureConfidence float32 `json:"time_signature_confidence"`
	} `json:"sections"`
	// Each segment contains a roughly conisistent sound throughout its duration.
	Segments []struct {
		// The starting point (in seconds) of the segment.
		Start float32 `json:"start"`
		// The duration (in seconds) of the segment.
		Duration float32 `json:"duration"`
		// The confidence, from 0.0 to 1.0, of the reliability of the segmentation. Segments of the song which are difficult to logically segment (e.g: noise) may correspond to low values in this field.
		Confidence float32 `json:"confidence"`
		// The onset loudness of the segment in decibels (dB). Combined with loudness_max and loudness_max_time, these components can be used to describe the "attack" of the segment.
		LoudnessStart float32 `json:"loudness_start"`
		// The peak loudness of the segment in decibels (dB). Combined with loudness_start and loudness_max_time, these components can be used to describe the "attack" of the segment.
		LoudnessMax float32 `json:"loudness_max"`
		// The segment-relative offset of the segment peak loudness in seconds. Combined with loudness_start and loudness_max, these components can be used to desctibe the "attack" of the segment.
		LoudnessMaxTime float32 `json:"loudness_max_time"`
		// The offset loudness of the segment in decibels (dB). This value should be equivalent to the loudness_start of the following segment.
		LoudnessEnd float32 `json:"loudness_end"`
		// Pitch content is given by a “chroma” vector, corresponding to the 12 pitch classes C, C#, D to B, with values ranging from 0 to 1 that describe the relative dominance of every pitch in the chromatic scale. For example a C Major chord would likely be represented by large values of C, E and G (i.e. classes 0, 4, and 7).
		//
		// Vectors are normalized to 1 by their strongest dimension, therefore noisy sounds are likely represented by values that are all close to 1, while pure tones are described by one value at 1 (the pitch) and others near 0. As can be seen below, the 12 vector indices are a combination of low-power spectrum values at their respective pitch frequencies.
		Pitches []float32 `json:"pitches"`
		// Timbre is the quality of a musical note or sound that distinguishes different types of musical instruments, or voices. It is a complex notion also referred to as sound color, texture, or tone quality, and is derived from the shape of a segment’s spectro-temporal surface, independently of pitch and loudness. The timbre feature is a vector that includes 12 unbounded values roughly centered around 0. Those values are high level abstractions of the spectral surface, ordered by degree of importance.
		//
		// For completeness however, the first dimension represents the average loudness of the segment; second emphasizes brightness; third is more closely correlated to the flatness of a sound; fourth to sounds with a stronger attack; etc. See an image below representing the 12 basis functions (i.e. template segments).
		//
		// The actual timbre of the segment is best described as a linear combination of these 12 basis functions weighted by the coefficient values: timbre = c1 x b1 + c2 x b2 + ... + c12 x b12, where c1 to c12 represent the 12 coefficients and b1 to b12 the 12 basis functions as displayed below. Timbre vectors are best used in comparison with each other.
		Timbre []float32 `json:"timbre"`
	} `json:"segments"`
	// A tatum represents the lowest regular pulse train that a listener intuitively infers from the timing of perceived musical events (segments).
	Tatums []struct {
		// The starting point (in seconds) of the time interval.
		Start float32 `json:"start"`
		// The duration (in seconds) of the time interval.
		Duration float32 `json:"duration"`
		// The confidence, from 0.0 to 1.0, of the reliability of the interval.
		Confidence float32 `json:"confidence"`
	} `json:"tatums"`
}

type GetTrackParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
}

type GetTrackResponse struct {
	Track
}

func (c *Client) GetTrack(id string, market string) (*GetTrackResponse, error) {
	track := GetTrackResponse{}
	var err *SpotifyError

	params := GetTrackParams{
		Market: market,
	}

	c.get(fmt.Sprintf("/tracks/%s", id)).QueryStruct(params).Receive(&track, err)

	if err != nil {
		return nil, err
	}

	return &track, nil
}

type GetSeveralTracksParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
	// A comma-separated list of the Spotify IDs.
	IDs string `url:"ids"`
}

type GetSeveralTracksResponse struct {
	Tracks []Track `json:"tracks"`
}

// Get Spotify catalog information for multiple tracks based on their Spotify IDs.
func (c *Client) GetSeveralTracks(market string, ids []string) (*GetSeveralTracksResponse, error) {
	tracks := GetSeveralTracksResponse{}
	var err *SpotifyError

	params := GetSeveralTracksParams{
		Market: market,
		IDs:    strings.Join(ids, ","),
	}

	c.get("/tracks").QueryStruct(params).Receive(&tracks, err)

	if err != nil {
		return nil, err
	}

	return &tracks, nil
}

type GetUsersSavedTracksParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit,omitempty"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset,omitempty"`
}

type GetUsersSavedTracksResponse struct {
	Pagination
	Items []SavedTrack `json:"items"`
}

// Get a list of the songs saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedTracks(market string, limit int, offset int) (*GetUsersSavedTracksResponse, error) {
	tracks := GetUsersSavedTracksResponse{}
	var err *SpotifyError

	params := GetUsersSavedTracksParams{
		Market: market,
		Limit:  limit,
		Offset: offset,
	}

	c.get("/me/tracks").QueryStruct(params).Receive(&tracks, err)

	if err != nil {
		return nil, err
	}

	return &tracks, nil
}

type SaveTracksForCurrentUserParams struct {
	// A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	IDs string `url:"ids"`
}

type SaveTracksForCurrentUserBody struct {
	// A JSON array of the Spotify IDs.
	//
	// A maximum of 50 items can be specified in one request. Note: if the ids parameter is present in the query string, any IDs listed here in the body will be ignored.
	IDs []string `json:"ids"`
}

type SaveTracksForCurrentUserResponse string

var saveTracksForCurrentUserResponse SaveTracksForCurrentUserResponse = "Track saved"

// Save one or more tracks to the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) SaveTracksForCurrentUser(ids []string) (*SaveTracksForCurrentUserResponse, error) {
	var err *SpotifyError

	params := SaveTracksForCurrentUserParams{
		IDs: strings.Join(ids, ","),
	}

	payload := SaveTracksForCurrentUserBody{
		IDs: ids,
	}

	c.put("/me/tracks").QueryStruct(params).BodyJSON(payload).Receive(&saveTracksForCurrentUserResponse, err)

	if err != nil {
		return nil, err
	}

	return &saveTracksForCurrentUserResponse, nil
}

type RemoveUsersSavedTracksParams struct {
	// A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	IDs string `url:"ids"`
}

type RemoveUsersSavedTracksBody struct {
	// A JSON array of the Spotify IDs.
	//
	// A maximum of 50 items can be specified in one request. Note: if the ids parameter is present in the query string, any IDs listed here in the body will be ignored.
	IDs []string `json:"ids"`
}

type RemoveUsersSavedTracksResponse string

var removeUsersSavedTracksResponse RemoveUsersSavedTracksResponse = "Track removed"

// Remove one or more tracks from the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedTracks(ids []string) (*RemoveUsersSavedTracksResponse, error) {
	var err *SpotifyError

	params := RemoveUsersSavedTracksParams{
		IDs: strings.Join(ids, ","),
	}

	payload := RemoveUsersSavedTracksBody{
		IDs: ids,
	}

	c.delete("/me/tracks").QueryStruct(params).BodyJSON(payload).Receive(removeUsersSavedTracksResponse, err)

	if err != nil {
		return nil, err
	}

	return &removeUsersSavedTracksResponse, nil
}

type CheckUsersSavedTracksParams struct {
	// A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	IDs string `url:"ids"`
}

type CheckUsersSavedTracksResponse []bool

// Check if one or more tracks is already saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedTracks(ids []string) (*CheckUsersSavedTracksResponse, error) {
	foundEach := CheckUsersSavedTracksResponse{}
	var err *SpotifyError

	params := CheckUsersSavedTracksParams{
		IDs: strings.Join(ids, ","),
	}

	c.get("/me/tracks/contains").QueryStruct(params).Receive(&foundEach, err)

	if err != nil {
		return nil, err
	}

	return &foundEach, nil
}

type GetMultiTracksAudioFeaturesParams struct {
	// A comma-separated list of the Spotify IDs for the tracks. Maximum: 100 IDs.
	IDs string `url:"ids"`
}

type GetMultiTracksAudioFeaturesResponse struct {
	AudioFeatures []AudioFeatures `json:"audio_features"`
}

// Get audio features for multiple tracks based on their Spotify IDs.
func (c *Client) GetMultiTracksAudioFeatures(ids []string) (*GetMultiTracksAudioFeaturesResponse, error) {
	features := GetMultiTracksAudioFeaturesResponse{}
	var err *SpotifyError

	params := GetMultiTracksAudioFeaturesParams{
		IDs: strings.Join(ids, ","),
	}

	c.get("/audio-features").QueryStruct(params).Receive(&features, err)

	if err != nil {
		return nil, err
	}

	return &features, nil
}

type GetSingleTracksAudioFeaturesResponse struct {
	AudioFeatures
}

// Get Track's Audio Features
//
// Get audio feature information for a single track identified by its unique Spotify ID.
func (c *Client) GetSingleTracksAudioFeatures(id string) (*GetSingleTracksAudioFeaturesResponse, error) {
	features := GetSingleTracksAudioFeaturesResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/audio-features/%s", id)).Receive(&features, err)

	if err != nil {
		return nil, err
	}

	return &features, nil
}

type GetTracksAudioAnalysisResponse struct {
	AudioAnalysis
}

// Get a low-level audio analysis for a track in the Spotify catalog. The audio analysis describes the track’s structure and musical content, including rhythm, pitch, and timbre.
func (c *Client) GetTracksAudioAnalysis(id string) (*GetTracksAudioAnalysisResponse, error) {
	analysis := GetTracksAudioAnalysisResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/audio-analysis/%s", id)).Receive(&analysis, err)

	if err != nil {
		return nil, err
	}

	return &analysis, nil
}

type GetRecommendationsBody struct {
	// The target size of the list of recommended tracks. For seeds with unusually small pools or when highly restrictive filtering is applied, it may be impossible to generate the requested number of recommended tracks. Debugging information for such cases is available in the response. Default: 20. Minimum: 1. Maximum: 100.
	Limit int `json:"limit,omitempty"`
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `json:"market,omitempty"`
	// A comma separated list of Spotify IDs for seed artists. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	//
	// Note: only required if seed_genres and seed_tracks are not set.
	SeedArtists string `json:"seed_artists"`
	// A comma separated list of any genres in the set of available genre seeds. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	//
	// Note: only required if seed_artists and seed_tracks are not set.
	SeedGenres string `json:"seed_genres"`
	// A comma separated list of Spotify IDs for a seed track. Up to 5 seed values may be provided in any combination of seed_artists, seed_tracks and seed_genres.
	//
	// Note: only required if seed_artists and seed_genres are not set.
	SeedTracks string `json:"seed_tracks"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinAcousticness float32 `json:"min_acousticness,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxAcousticness float32 `json:"max_acousticness,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetAcousticness float32 `json:"target_acousticness,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinDanceability float32 `json:"min_danceability,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxDanceability float32 `json:"max_danceability,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetDanceability float32 `json:"target_danceability,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinDurationMS int `json:"min_duration_ms,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxDurationMS int `json:"max_duration_ms,omitempty"`
	// Target duration of the track (ms)
	TargetDurationMS int `json:"target_duration_ms,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinEnergy float32 `json:"min_energy,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxEnergy float32 `json:"max_energy,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetEnergy float32 `json:"target_energy,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinInstrumentalness float32 `json:"min_instrumentalness,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxInstrumentalness float32 `json:"max_instrumentalness,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetInstrumentalness float32 `json:"target_instrumentalness,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinKey int `json:"min_key,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxKey int `json:"max_key,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetKey int `json:"target_key,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinLiveness float32 `json:"min_liveness,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxLiveness float32 `json:"max_liveness,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetLiveness float32 `json:"target_liveness,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinLoudness float32 `json:"min_loudness,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxLoudness float32 `json:"max_loudness,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetLoudness float32 `json:"target_loudness,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinMode int `json:"min_mode,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxMode int `json:"max_mode,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetMode int `json:"target_mode,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinPopularity int `json:"min_popularity,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxPopularity int `json:"max_popularity,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetPopularity int `json:"target_popularity,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinSpeechiness float32 `json:"min_speechiness,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxSpeechiness float32 `json:"max_speechiness,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetSpeechiness float32 `json:"target_speechiness,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinTempo float32 `json:"min_tempo,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxTempo float32 `json:"max_tempo,omitempty"`
	// Target tempo (BPM)
	TargetTempo float32 `json:"target_tempo,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinTimeSignature int `json:"min_time_signature,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxTimeSignature int `json:"max_time_signature,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetTimeSignature int `json:"target_time_signature,omitempty"`
	// For each tunable track attribute, a hard floor on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, min_tempo=140 would restrict results to only those tracks with a tempo of greater than 140 beats per minute.
	MinValence float32 `json:"min_valence,omitempty"`
	// For each tunable track attribute, a hard ceiling on the selected track attribute’s value can be provided. See tunable track attributes below for the list of available options. For example, max_instrumentalness=0.35 would filter out most tracks that are likely to be instrumental.
	MaxValence float32 `json:"max_valence,omitempty"`
	// For each of the tunable track attributes (below) a target value may be provided. Tracks with the attribute values nearest to the target values will be preferred. For example, you might request target_energy=0.6 and target_danceability=0.8. All target values will be weighed equally in ranking results.
	TargetValence float32 `json:"target_valence,omitempty"`
}

type GetRecommendationsResponse struct {
	// An array of recommendation seed objects.
	Seeds []struct {
		// The number of tracks available after min_* and max_* filters have been applied.
		AfterFilteringSize int `json:"afterFilteringSize"`
		// The number of tracks available after relinking for regional availability.
		AfterRelinkingSize int `json:"afterRelinkingSize"`
		// A link to the full track or artist data for this seed. For tracks this will be a link to a Track Object. For artists a link to an Artist Object. For genre seeds, this value will be null.
		Href string `json:"href"`
		// The id used to select this seed. This will be the same as the string used in the seed_artists, seed_tracks or seed_genres parameter.
		ID string `json:"id"`
		// The number of recommended tracks available for this seed.
		InitialPoolSize int `json:"initialPoolSize"`
		// The entity type of this seed. One of artist, track or genre.
		Type string `json:"type"`
	} `json:"seeds"`
	// An array of track object (simplified) ordered according to the parameters supplied.
	Tracks []Track `json:"tracks"`
}

// Recommendations are generated based on the available information for a given seed entity and matched against similar artists and tracks. If there is sufficient information about the provided seeds, a list of tracks will be returned together with pool size details.
//
// For artists and tracks that are very new or obscure there might not be enough data to generate a list of tracks.
func (c *Client) GetRecommendations(payload GetRecommendationsBody) (*GetRecommendationsResponse, error) {
	recs := GetRecommendationsResponse{}
	var err *SpotifyError

	c.get("/recommendations").BodyJSON(payload).Receive(&recs, err)

	if err != nil {
		return nil, err
	}

	return &recs, nil
}
