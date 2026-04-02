package handlers

import (
	"encoding/json"
	"net/http"

	"permit-proxy/internal/store"
)

type PermitHandler struct {
	store *store.Store
}

func (h *PermitHandler) HandlePermits(w http.ResponseWriter, r *http.Request) {
	g, err := h.store.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func NewPermitHandler(s *store.Store) *PermitHandler {
	p := &PermitHandler{
		store: s,
	}

	return p
}
