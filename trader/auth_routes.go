package main

import (
	"encoding/json"
	"net/http"

	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

type loginRequest struct {
	email    string
	password string
}

func (t *trader) login(w http.ResponseWriter, r *http.Request) {
	var body loginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.InternalServerError(w, err)
		return
	}
	user, err := traderdb.GetUserWithEmail(t.db, body.email)
	if err != nil {
		t.logger.Error("Failed to GetUserWithEmail:", err)
		utils.UnauthenticatedError(w, invalidCredentialsError)
		return
	}
	if authenticated := utils.CompareHashToPassword(
		user.HashedPassword,
		body.password,
	); !authenticated {
		utils.UnauthenticatedError(w, invalidCredentialsError)
		return
	}

	session := t.getSession(r)
	session.Values["userID"] = user.ID
	if err := session.Save(r, w); err != nil {
		t.logger.Error("Failed to save session:", err)
	}
}

func (t *trader) logout(w http.ResponseWriter, r *http.Request) {
	session := t.getSession(r)
	delete(session.Values, "userID")
	if err := session.Save(r, w); err != nil {
		t.logger.Error("Failed to save session:", err)
	}
}

func (t *trader) addAuthRoutes(router chi.Router) {
	router.Post("/login", t.login)
	router.Post("/logout", t.logout)
}

func (t *trader) getSessionUserID(r *http.Request) (int, bool) {
	session := t.getSession(r)
	userID, ok := session.Values["userID"].(int)
	if userID == 0 || !ok {
		return 0, false
	}
	return userID, true
}

func (t *trader) getSession(r *http.Request) *sessions.Session {
	session, err := t.sessionStore.Get(r, "trader-session")
	if err != nil {
		t.logger.Error("Failed to get session:", err)
	}
	return session
}
