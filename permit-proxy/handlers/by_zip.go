package handlers

import (
	"encoding/json"
	"net/http"

	"permit-proxy/internal/aggregator"
)

func (h *PermitHandler) HandleByZip(w http.ResponseWriter, r *http.Request) {
	g, err := h.store.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a, err := aggregator.ByZip(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}
