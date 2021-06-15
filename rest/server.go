package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "PONG")
}

func small(w http.ResponseWriter, req *http.Request) {
	var reqBody smallRequest
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		writeErr(w, http.StatusInternalServerError)
		return
	}

	username := reqBody.Identity.Username
	email := reqBody.Identity.Email
	country := reqBody.Identity.Country

	respBody := smallResponse{
		Summary: fmt.Sprintf("username=%s email=%s country=%s", username, email, country),
	}

	if err := json.NewEncoder(w).Encode(&respBody); err != nil {
		writeErr(w, http.StatusInternalServerError)
	}
}

func big(w http.ResponseWriter, req *http.Request) {
	var reqBody bigRequest
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		writeErr(w, http.StatusInternalServerError)
		return
	}

	username := reqBody.Identity.Username
	email := reqBody.Identity.Email
	country := reqBody.Identity.Country
	lat := reqBody.Location.Lat
	lon := reqBody.Location.Lon
	interest := reqBody.Interest
	bookmark := reqBody.Bookmark

	bookmarkedInterest := make(map[string]bool, len(interest))
	organizedBookmark := make(map[string][]*content)

	for _, bm := range bookmark {
		var cl []*content
		if v, ok := organizedBookmark[bm.Genre]; !ok {
			cl = []*content{}
		} else {
			cl = v
		}

		organizedBookmark[bm.Genre] = append(cl, bm)
	}

	for _, v := range interest {
		_, ok := organizedBookmark[v]
		bookmarkedInterest[v] = ok
	}

	respBody := bigResponse{
		Summary: fmt.Sprintf("username=%s email=%s country=%s lat=%f lon=%f", username, email, country, lat, lon),
		BookmarkedInterest: bookmarkedInterest,
		OrganizedBookmark: organizedBookmark,
	}

	if err := json.NewEncoder(w).Encode(&respBody); err != nil {
		writeErr(w, http.StatusInternalServerError)
	}
}

func writeErr(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&errorResponse{
		Message: "something went wrong",
	})
}