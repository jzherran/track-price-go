package register

import (
	"context"
	"net/http"

	"github.com/jzherran/track-price-go/pkg/http/render"
	"github.com/jzherran/track-price-go/trackitem"
)

type (
	// TrackItemManager interface to manage data about track item handler.
	TrackItemManager interface {
		FindAllTrackItems(context.Context) (trackitem.TrackItems, error)
	}

	// TrackItemHandler handler to use in terms of manage track items.
	TrackItemHandler struct {
		manager TrackItemManager
	}
)

// NewTrackItemHandler return a new instance of the handler.
func NewTrackItemHandler(m TrackItemManager) TrackItemHandler {
	return TrackItemHandler{
		manager: m,
	}
}

func (h TrackItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	trackItems, _ := h.manager.FindAllTrackItems(r.Context())
	render.JSON(w, http.StatusOK, trackItems)
}
