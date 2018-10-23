package collector

import (
	"context"
	"sync"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger/microloggertest"
	"github.com/prometheus/client_golang/prometheus"
)

func Test_Collector_Set_Boot(t *testing.T) {
	for i := 0; i < 100; i++ {
		c, s, err := newCollectorAndSet()
		if err != nil {
			t.Fatalf(err.Error())
		}

		var wg sync.WaitGroup

		for g := 0; g < 100; g++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				err := s.Boot(context.Background())
				if err != nil {
					t.Fatalf(err.Error())
				}
			}()
		}

		wg.Wait()

		if !c.gotBootedOnce() {
			t.Fatalf("expected collector to be booted once")
		}
	}
}

//
//
//

func newCollectorAndSet() (*bootCountCollector, *Set, error) {
	var err error

	var newCollector *bootCountCollector
	{
		newCollector = &bootCountCollector{}
	}

	var newSet *Set
	{
		c := SetConfig{
			Collectors: []Interface{
				newCollector,
			},
			Logger: microloggertest.New(),
		}

		newSet, err = NewSet(c)
		if err != nil {
			return nil, nil, microerror.Mask(err)
		}
	}

	return newCollector, newSet, nil
}

//
//
//

type bootCountCollector struct {
	bootCount int
	mutex     sync.Mutex
}

func (c *bootCountCollector) Boot(ctx context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.bootCount++

	return nil
}

func (c *bootCountCollector) CollectWithError(ch chan<- prometheus.Metric) error {
	return nil
}

func (c *bootCountCollector) DescribeWithError(ch chan<- *prometheus.Desc) error {
	return nil
}

func (c *bootCountCollector) gotBootedOnce() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.bootCount == 1
}
