package torrent

// Very basic torrent info storage
type Torrent struct {
	Announce string
	Info     map[string]any
}
