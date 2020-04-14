package cachetype

import (
	"time"

	"github.com/hashicorp/consul/agent/cache"
)

type RegisterOptionsBlockingRefresh struct{}

func (r RegisterOptionsBlockingRefresh) RegisterOptions() cache.RegisterOptions {
	return cache.RegisterOptions{
		// Maintain a blocking query, retry dropped connections quickly
		Refresh:          true,
		SupportsBlocking: true,
		RefreshTimer:     0 * time.Second,
		RefreshTimeout:   10 * time.Minute,
	}
}

type RegisterOptionsNoRefresh struct{}

func (r RegisterOptionsNoRefresh) RegisterOptions() cache.RegisterOptions {
	return cache.RegisterOptions{
		Refresh:          false,
		SupportsBlocking: false,
	}
}
