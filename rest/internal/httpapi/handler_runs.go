package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/SvBrunner/there-and-back-again/internal/service"
)

type RunsHandler struct {
	svc service.Service
}

func NewRunsHandler(svc service.Service) *RunsHandler {
	return &RunsHandler{svc: svc}
}

type addRunRequest struct {
	DistanceKm    float64 `json:"distance_km"`
	TimeInMinutes int32   `json:"time_in_minutes"`
	Date          string  `json:"date,omitempty"`
}

func (h *RunsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.addRun(w, r)
	case http.MethodGet:
		h.listRuns(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Use GET or POST.")
	}
}

func (h *RunsHandler) addRun(w http.ResponseWriter, r *http.Request) {
	var req addRunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "Body must be valid JSON.")
		return
	}

	if req.DistanceKm <= 0 {
		writeError(w, http.StatusBadRequest, "invalid_distance", "distance_km must be > 0.")
		return
	}

	if req.Date != "" {
		if _, err := time.Parse(time.RFC3339, req.Date); err != nil {
			writeError(w, http.StatusBadRequest, "invalid_date", "date must be RFC3339 (e.g. 2026-01-18T10:00:00Z).")
			return
		}
	}

	run, err := h.svc.AddRun(r.Context(), req.DistanceKm, req.TimeInMinutes)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, "not_found", "Resource not found.")
			return
		}
		writeError(w, http.StatusInternalServerError, "internal_error", "Something went wrong.")
		return
	}

	writeJSON(w, http.StatusCreated, run)
}

func (h *RunsHandler) listRuns(w http.ResponseWriter, r *http.Request) {
	runs, err := h.svc.ListRuns(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal_error", "Something went wrong.")
		return
	}

	writeJSON(w, http.StatusOK, runs)
}
