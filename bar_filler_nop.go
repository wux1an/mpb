package mpb

import (
	"io"

	"github.com/wux1an/mpb/v7/decor"
)

// NopStyle provides BarFillerBuilder which builds NOP BarFiller.
func NopStyle() BarFillerBuilder {
	return BarFillerBuilderFunc(func() BarFiller {
		return BarFillerFunc(func(io.Writer, int, decor.Statistics) {})
	})
}
