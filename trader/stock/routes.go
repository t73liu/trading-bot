package stock

import (
	"log"
)

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) AddRoutes() {
	// TODO fetch all stocks, single stock bars, ...
}
