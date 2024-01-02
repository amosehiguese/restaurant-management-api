package handlers

import (
	"context"

	"github.com/amosehiguese/restaurant-api/log"
)

type resp map[string]any

var l = log.NewLog()

var ctx = context.Background()
