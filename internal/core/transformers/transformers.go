package transformers

import (
	"github.com/Zanda256/ike-go/internal/core/transformers/wpTransform"
	"github.com/Zanda256/ike-go/pkg-foundation/logger"
)

type TransformService struct {
	log    logger.Logger
	WPress wpTransform.WpTransformer
}

func (ts *TransformService) TrasformWp(hosts []string) {
	err := ts.WPress.Transform(hosts)
}
