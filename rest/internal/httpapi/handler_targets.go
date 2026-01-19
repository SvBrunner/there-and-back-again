package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/SvBrunner/there-and-back-again/internal/service"
)

type TargetsHandler struct {
	svc service.Service
}

func NewTargetsHandler(svc service.Service) *TargetsHandler {
	return &TargetsHandler{svc: svc}
}

type addTargetRequest struct {
	Distance float64 `json:"distance"`
	Name     string  `json:"name"`
}

func (h *TargetsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.addTarget(w, r)
	case http.MethodGet:
		h.listTargets(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Use GET or POST.")
	}
}

func (h *TargetsHandler) addTarget(w http.ResponseWriter, r *http.Request) {
	var req addTargetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "Body must be valid JSON.")
		return
	}

	if req.Distance <= 0 {
		writeError(w, http.StatusBadRequest, "invalid_distance", "distance must be > 0")
		return
	}

	if len(req.Name) <= 0 {
		writeError(w, http.StatusBadRequest, "invalid_name", "name must be set")
	}
	run, err := h.svc.AddTarget(r.Context(), req.Distance, req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal_error", "Something went wrong.")
		return
	}

	writeJSON(w, http.StatusCreated, run)
}

func (h *TargetsHandler) listTargets(w http.ResponseWriter, r *http.Request) {
	runs, err := h.svc.ListTargets(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal_error", "Something went wrong.")
		return
	}
	writeJSON(w, http.StatusOK, runs)
}
