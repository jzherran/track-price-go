package register

import (
	"context"
	"fmt"

	"github.com/jzherran/track-price-go/internal"
	"github.com/jzherran/track-price-go/trackitem"
)

type (
	repository struct {
		datasource internal.Datasource
	}
)

// NewTrackItemRetriverRepository instance of a repository.
func NewTrackItemRetriverRepository(ds internal.Datasource) TrackItemManager {
	return repository{
		datasource: ds,
	}
}

func (r repository) FindAllTrackItems(ctx context.Context) (trackitem.TrackItems, error) {
	result := []trackitem.TrackItem{}
	result = append(result, trackitem.TrackItem{
		Name:    "[alkosto::nevera]",
		URLLink: "https://www.alkosto.com/nevera-samsung-394-litros-negro-rt38k5992bs",
	})

	values := r.datasource.Read("Sheet1!A1:E")

	if len(values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range values {
			for _, col := range row {
				fmt.Printf("%s ", col)
			}
			fmt.Println()
		}
	}

	return result, nil
}
