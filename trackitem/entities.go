package trackitem

// TrackItem is the object used to get the all data registered in the spreedsheat configured to get
// the list of items related to the process
type TrackItem struct {
	Name    string `json:"name"`
	URLLink string `json:"urlLink"`
}

// TrackItems list of track items
type TrackItems []TrackItem
