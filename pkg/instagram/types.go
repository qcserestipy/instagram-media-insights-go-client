package instagram

// Reel represents an Instagram reel with its metrics
type Reel struct {
	ID              string  `json:"id"`
	DateTime        string  `json:"date_time"`
	Views           int     `json:"views"`
	Reach           int     `json:"reach"`
	Likes           int     `json:"likes"`
	Comments        int     `json:"comments"`
	Shares          int     `json:"shares"`
	Saves           int     `json:"saves"`
	EngagementViews float64 `json:"engagement_views"`
	Caption         string  `json:"caption"`
	// EngagementFollowers float64 `json:"engagement_followers"`
}
