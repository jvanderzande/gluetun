package updater

import (
	"time"

	"github.com/qdm12/gluetun/internal/updater/resolver"
)

func newParallelResolver() (parallelResolver *resolver.Parallel) {
	const (
		maxFailRatio    = 1
		maxDuration     = 20 * time.Second
		betweenDuration = time.Second
		maxNoNew        = 4
		maxFails        = 10
	)
	settings := resolver.ParallelSettings{
		MaxFailRatio: maxFailRatio,
		Repeat: resolver.RepeatSettings{
			MaxDuration:     maxDuration,
			BetweenDuration: betweenDuration,
			MaxNoNew:        maxNoNew,
			MaxFails:        maxFails,
			SortIPs:         true,
		},
	}
	return resolver.NewParallelResolver(settings)
}