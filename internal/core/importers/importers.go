package importers

import (
	"github.com/Zanda256/ike-go/internal/core/importers/wpImport"
)

type ImportService struct {
	WPress *wpImport.ImportManager
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
