package bluelink

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
)

const (
	KiaAppID     = "e7bcd186-a5fd-410d-92cb-6876a42288bd"
	HyundaiAppID = "014d2225-8495-4735-812d-2616334fd15d"
)

// StampsRegistry collects stamps for a single brand
type StampsRegistry map[string]*StampCollection

type StampCollection struct {
	Stamps    []string
	Generated time.Time
	Frequency float64
}

var (
	client = request.NewHelper(util.NewLogger("http"))

	brands = map[string]string{
		KiaAppID:     "kia",
		HyundaiAppID: "hyundai",
	}

	mu     sync.Mutex
	Stamps = StampsRegistry{
		KiaAppID:     nil,
		HyundaiAppID: nil,
	}
)

func download(log *util.Logger, id, brand string) error {
	var res StampCollection
	uri := fmt.Sprintf("https://raw.githubusercontent.com/neoPix/bluelinky-stamps/master/%s-%s.v2.json", brand, id)

	if err := client.GetJSON(uri, &res); err != nil {
		return fmt.Errorf("failed to download stamps: %w", err)
	}

	mu.Lock()
	Stamps[id] = &res
	mu.Unlock()

	return nil
}

// updateStamps updates stamps according to https://github.com/Hacksore/bluelinky/pull/144
func updateStamps(log *util.Logger, id string) error {
	mu.Lock()
	if Stamps[id] != nil {
		mu.Unlock()
		return nil
	}
	mu.Unlock()

	if err := download(log, id, brands[id]); err != nil {
		return err
	}

	go func() {
		for range time.NewTicker(12 * time.Hour).C {
			if err := download(log, id, brands[id]); err != nil {
				log.ERROR.Println(err)
			}
		}
	}()

	return nil
}

// New creates a new stamp
func (s StampsRegistry) New(id string) string {
	mu.Lock()
	defer mu.Unlock()

	source := s[id]
	if source == nil {
		panic(id)
	}

	position := float64(time.Since(source.Generated).Milliseconds()) / source.Frequency

	return source.Stamps[int64(position+5*rand.Float64())]
}
