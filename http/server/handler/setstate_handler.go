package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

type SetStateRequest struct {
	Fill int `json:"fill"`
}

func (ls *LifeStates) setState(w http.ResponseWriter, r *http.Request) {
	var req SetStateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ls.logger.Error("Invalid request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Fill < 0 || req.Fill > 100 {
		ls.logger.Error("Invalid fill percentage")
		http.Error(w, "Invalid fill percentage", http.StatusBadRequest)
		return
	}

	// Save the fill percentage to state.cfg
	err := os.WriteFile("state.cfg", []byte(fmt.Sprintf("%d%%", req.Fill)), 0644)
	if err != nil {
		ls.logger.Error("Failed to save state", zap.Error(err))
		http.Error(w, "Failed to save state", http.StatusInternalServerError)
		return
	}

	ls.logger.Info("State successfully updated", zap.Int("fill", req.Fill))
	w.WriteHeader(http.StatusOK)
}

func (ls *LifeStates) resetState(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("state.cfg")
	if err != nil {
		ls.logger.Error("Failed to read state file", zap.Error(err))
		http.Error(w, "Failed to read state file", http.StatusInternalServerError)
		return
	}

	var fill int
	_, err = fmt.Sscanf(string(data), "%d%%", &fill)
	if err != nil {
		ls.logger.Error("Failed to parse state file", zap.Error(err))
		http.Error(w, "Failed to parse state file", http.StatusInternalServerError)
		return
	}

	response := map[string]int{"fill": fill}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		ls.logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	ls.logger.Info("State successfully reset", zap.Int("fill", fill))
}
