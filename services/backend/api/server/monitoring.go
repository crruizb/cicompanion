package server

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/cicompanion/data"
)

type MonitoringURLDTO struct {
	Urls []data.MonitoringURL `json:"urls"`
}

func (rt *Router) addMonitoring(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	mURLs := &MonitoringURLDTO{}
	err := ReadJSON(w, r, mURLs)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	err = rt.ms.AddMonitoringURLs(mURLs.Urls, user.Id)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (rt *Router) getMonitoringUrls(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	mURLsDto := &MonitoringURLDTO{}

	mURLS, err := rt.ms.GetMonitoringURLsByUserId(user.Id)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	mURLsDto.Urls = mURLS
	WriteJSON(w, http.StatusOK, mURLsDto, nil)
}

func (rt *Router) getAllMonitoringUrlsToPing(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	if user.Role != "ADMIN" {
		BadRequestResponse(w, r, errors.New("user is not ADMIN"))
		return
	}
	mURLsDto := &MonitoringURLDTO{}

	mURLS, err := rt.ms.GetAllMonitoringURLs()
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	mURLsDto.Urls = mURLS
	WriteJSON(w, http.StatusOK, mURLsDto, nil)
}

func (rt *Router) updateMonitoringURL(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	mURLIdStr := r.PathValue("id")
	mURLId, err := strconv.Atoi(mURLIdStr)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	mURL := &data.MonitoringURL{}
	err = ReadJSON(w, r, mURL)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	updatedMURL, err := rt.ms.UpdateMonitoringURL(mURLId, user.Id, *mURL)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			BadRequestResponse(w, r, errors.New("monitoring url not found"))
			return
		}
		ServerErrorResponse(w, r, err)
		return
	}

	WriteJSON(w, http.StatusOK, updatedMURL, nil)
}

func (rt *Router) deleteMonitoringURL(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(ContextUser).(*data.User)
	mURLIdStr := r.PathValue("id")
	mURLId, err := strconv.Atoi(mURLIdStr)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

	err = rt.ms.DeleteMonitoringURL(mURLId, user.Id)
	if err != nil {
		ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
