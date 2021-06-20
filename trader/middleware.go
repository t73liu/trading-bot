package main

import (
	"context"
	"net/http"
	"time"

	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"
)

func (t *trader) logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		t.logger.Infof("Starting: %s - %s\n", r.Method, r.URL)
		defer func() {
			t.logger.Infof(
				"Completed (%dms): %s - %s\n",
				time.Since(startTime).Milliseconds(),
				r.Method,
				r.URL.Path,
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func (t *trader) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, authenticated := t.getSessionUserID(r)
		if !authenticated {
			utils.UnauthenticatedError(w, unauthenticatedError)
			return
		}

		user, err := traderdb.GetUserWithID(t.db, userID)
		if err != nil || !user.IsActive {
			t.logout(w, r)
			utils.UnauthenticatedError(w, unauthenticatedError)
			return
		}

		ctx := context.WithValue(r.Context(), userIDContextKey, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getContextUserID returns the userID stored in request.context. This
// helper should only be called behind the requireAuthentication middleware
func getContextUserID(r *http.Request) int {
	userID, ok := r.Context().Value(userIDContextKey).(int)
	if !ok {
		return 0
	}
	return userID
}
