package templatedata

import "github.com/EngineerKamesh/igb/igweb/shared/models"

type About struct {
	PageTitle string
	Gophers   []*models.Gopher
}
