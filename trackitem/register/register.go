package register

import (
	"github.com/go-chi/chi"
	"github.com/jzherran/track-price-go/internal"
)

// SetupModule configure and start the module to register track items
func SetupModule(router chi.Router, ds internal.Datasource) {
	repo := NewTrackItemRetriverRepository(ds)

	handler := NewTrackItemHandler(repo)

	router.Method("GET", "/", handler)
}
