package mp

type PlaybackStatus string

const (
	// -- PLAYBACK STATUS --
	// Support
	// - Windows
	// - Linux
	// - macOS (* needs to be built as proxy using AVPlayer.rate == 0)
	PlaybackStatusPlaying = "Playing"
	PlaybackStatusPaused  = "Paused"
	// Support
	// - Windows
	// - Linux
	PlaybackStatusStopped = "Stopped"
	// Support
	// - Windows
	PlaybackStatusChanging = "Changing"
	PlaybackStatusClosed   = "Closed"
	PlaybackStatusError    = "PlaybackStatusError"
)
