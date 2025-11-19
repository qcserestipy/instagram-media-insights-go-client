package instagram

import "time"

// Reel represents an Instagram reel with its metrics
type Reel struct {
	ID                string    `json:"id"`
	DateTime          time.Time `json:"date_time"`
	Views             int64     `json:"views"`
	Reach             int64     `json:"reach"`
	Likes             int64     `json:"likes"`
	Comments          int64     `json:"comments"`
	Shares            int64     `json:"shares"`
	Saves             int64     `json:"saves"`
	TotalInteractions int64     `json:"total_interactions"`
	EngagementViews   float64   `json:"engagement_views"`
	Caption           string    `json:"caption"`
	// EngagementFollowers float64 `json:"engagement_followers"`
}

// Story represents an Instagram story with its metrics
type Story struct {
	ID                string    `json:"id"`
	DateTime          time.Time `json:"date_time"`
	Views             int64     `json:"views"`
	Reach             int64     `json:"reach"`
	Replies           int64     `json:"replies"`
	Shares            int64     `json:"shares"`
	Navigation        int64     `json:"navigation"`
	ProfileActivity   int64     `json:"profile_activity"`
	ProfileVisits     int64     `json:"profile_visits"`
	Follows           int64     `json:"follows"`
	TotalInteractions int64     `json:"total_interactions"`
	EngagementViews   float64   `json:"engagement_views"`
}

type Demographics struct {
	AgeRanges map[string]int64 `json:"age_ranges"`
	Genders   map[string]int64 `json:"genders"`
	Countries map[string]int64 `json:"countries"`
	Cities    map[string]int64 `json:"cities"`
}

type AccountDemographics struct {
	Follower *Demographics `json:"follower"`
	Engaged  *Demographics `json:"engaged"`
}

type AccountInfo struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	FollowersCount    int64  `json:"followers_count"`
	FollowingCount    int64  `json:"following_count"`
	MediaCount        int64  `json:"media_count"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Bio               string `json:"bio"`
	Website           string `json:"website"`
}

type FollowerDynamics struct {
	NewFollowers int64     `json:"new_followers"`
	Unfollowers  int64     `json:"unfollowers"`
	NetFollowers int64     `json:"net_followers"`
	Since        time.Time `json:"since"`
	Until        time.Time `json:"until"`
}
