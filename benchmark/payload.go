package main

type identity struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Country string `json:"country"`
}

type location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type content struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReactCount int64 `json:"react_count"`
	ShareCount int64 `json:"share_count"`
	Genre string `json:"genre"`
}

type smallRequest struct {
	Identity identity `json:"identity"`
}

type smallResponse struct {
	Summary string `json:"summary"`
}

type bigRequest struct {
	Identity identity `json:"identity"`
	Location location `json:"location"`
	Interest []string `json:"interest"`
	Bookmark []*content `json:"bookmark"`
}

type bigResponse struct {
	Summary string `json:"summary"`
	BookmarkedInterest map[string]bool `json:"bookmarked_interest"`
	OrganizedBookmark map[string][]*content `json:"organized_bookmark"`
}

type errorResponse struct {
	Message string `json:"message"`
}
