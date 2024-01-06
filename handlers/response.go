package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amosehiguese/restaurant-api/log"
)

type resp map[string]any

var l = log.NewLog()

var ctx = context.Background()

const pageSize = 4

func  paginate(w http.ResponseWriter, r *http.Request) (*int, *int, error) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageNum, err:= strconv.Atoi(page)
	if err != nil {
		return  nil, nil, err
	}
	start := (pageNum - 1) * pageSize
	end := 	start + pageSize
	return &start, &end, nil
}

func parseDate(dateString string) time.Time {
	parsedTime, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		l.Log.Fatal(err.Error())
	}
	return parsedTime
}


func parseTime(timeString string) time.Time {
	parsedTime, err := time.Parse("15:04:05", timeString)
	if err != nil {
		l.Log.Fatal(err.Error())
	}
	return parsedTime
}