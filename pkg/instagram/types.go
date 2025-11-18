package instagram

// Reel represents an Instagram reel with its metrics
type Reel struct {
	ID                string  `json:"id"`
	DateTime          string  `json:"date_time"`
	Views             int     `json:"views"`
	Reach             int     `json:"reach"`
	Likes             int     `json:"likes"`
	Comments          int     `json:"comments"`
	Shares            int     `json:"shares"`
	Saves             int     `json:"saves"`
	TotalInteractions int     `json:"total_interactions"`
	EngagementViews   float64 `json:"engagement_views"`
	Caption           string  `json:"caption"`
	// EngagementFollowers float64 `json:"engagement_followers"`
}

// Story represents an Instagram story with its metrics
type Story struct {
	ID                string  `json:"id"`
	DateTime          string  `json:"date_time"`
	Views             int     `json:"views"`
	Reach             int     `json:"reach"`
	Replies           int     `json:"replies"`
	Shares            int     `json:"shares"`
	Navigation        int     `json:"navigation"`
	ProfileActivity   int     `json:"profile_activity"`
	ProfileVisits     int     `json:"profile_visits"`
	Follows           int     `json:"follows"`
	TotalInteractions int     `json:"total_interactions"`
	EngagementViews   float64 `json:"engagement_views"`
}

type Demographics struct {
	AgeRanges map[string]int `json:"age_ranges"`
	Genders   map[string]int `json:"genders"`
	Countries map[string]int `json:"countries"`
	Cities    map[string]int `json:"cities"`
}

type AccountDemographics struct {
	Follower *Demographics `json:"follower"`
	Engaged  *Demographics `json:"engaged"`
}
