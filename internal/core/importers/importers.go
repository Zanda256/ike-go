package importers

import (
	"github.com/Zanda256/ike-go/internal/core/importers/wp"
)

type ImportService struct {
	WPress *wp.ImportManager
	// Add import managers here
}

func (s *ImportService) ImportWP(urls []string) error {
	for _, url := range urls {
		err := s.WPress.Import(url)
		if err != nil {
			return err
		}
	}
	return nil
}
