package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

type ctxKeyReqID int

const RequestIDKey ctxKeyReqID = iota

var (
	prefix string
	reqid uint64
)

func init() {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}

	var buf [12]byte
	var b64 string

	for len(b64) < 10 {
		rand.Read(buf[:])
		b64 = base64.StdEncoding.EncodeToString(buf[:])
		b64 = strings.NewReplacer("+", "", "/", "").Replace(b64)
	}

	prefix = fmt.Sprintf("%s/%s", hostname, b64[0:10])
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			reqID := r.Header.Get("X-Request-Id")
			if reqID == "" {
				myid := atomic.AddUint64(&reqid, 1)
				reqID = fmt.Sprintf("%s-%06d", prefix, myid)
			}
			ctx = context.WithValue(ctx, RequestIDKey, reqID)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					if err == http.ErrAbortHandler {
						panic(err)
					}
				}
			}()
			
			next.ServeHTTP(w,r)
		},
	)
}