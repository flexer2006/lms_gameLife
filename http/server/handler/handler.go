package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/flexer2006/lms_gameLife/internal/service"
	"go.uber.org/zap"
)

// Decorator type for adding middleware to handlers
type Decorator func(http.Handler) http.Handler

// LifeStates object for storing game state
type LifeStates struct {
	service.GameService
	logger *zap.Logger
}

func New(lifeService *service.GameService, logger *zap.Logger) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := &LifeStates{
		GameService: *lifeService,
		logger:      logger,
	}

	serveMux.HandleFunc("/next-state", lifeState.nextState)
	serveMux.HandleFunc("/setstate", lifeState.setState)
	serveMux.HandleFunc("/reset", lifeState.resetState) // Новый маршрут

	return serveMux, nil
}

// Decorate adds middleware to the handlers
func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}

	return decorated
}

// nextState gets the next game state
func (ls *LifeStates) nextState(w http.ResponseWriter, _ *http.Request) {
	worldState := ls.GameService.GetNextState()

	err := json.NewEncoder(w).Encode(worldState)
	if err != nil {
		ls.logger.Error("Game state coding error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ls.logger.Info("Game status successfully submitted")
}
