package handlers

import (
	"context"
	"net/http"
	"strconv"

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