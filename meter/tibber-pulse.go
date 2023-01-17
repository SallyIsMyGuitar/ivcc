package meter

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/meter/tibber"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/hasura/go-graphql-client"
)

func init() {
	registry.Add("tibber-pulse", NewTibberFromConfig)
}

var timeout = time.Minute

type Tibber struct {
	mu      sync.Mutex
	log     *util.Logger
	updated time.Time
	live    tibber.LiveMeasurement
}

func NewTibberFromConfig(other map[string]interface{}) (api.Meter, error) {
	var cc struct {
		Token  string
		HomeID string
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Token == "" {
		return nil, errors.New("missing token")
	}

	t := &Tibber{
		log: util.NewLogger("pulse").Redact(cc.Token, cc.HomeID),
	}

	// query client
	qclient := tibber.NewClient(t.log, cc.Token)

	if cc.HomeID == "" {
		if home, err := qclient.DefaultHome(""); err != nil {
			return nil, err
		} else {
			cc.HomeID = home.ID
		}
	}

	var res struct {
		Viewer struct {
			WebsocketSubscriptionUrl string
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), request.Timeout)
	defer cancel()

	if err := qclient.Query(ctx, &res, nil); err != nil {
		return nil, err
	}

	// subscription client
	// client := graphql.NewSubscriptionClient(tibber.SubscriptionURI).
	// WithProtocol(graphql.SubscriptionsTransportWS).
	// client := graphql.NewSubscriptionClient(res.Viewer.WebsocketSubscriptionUrl).
	client := graphql.NewSubscriptionClient("wss://websocket-api.tibber.com/v1-beta/gql/subscriptions").
		WithProtocol(graphql.GraphQLWS).
		WithConnectionParams(map[string]any{
			"token": cc.Token,
		}).
		WithRetryTimeout(timeout)
	// WithLog(t.log.TRACE.Println)

	// run the client
	done := make(chan error)
	go t.subscribe(client, cc.HomeID, done)
	err := <-done

	return t, err
}

// subscribe to the websocket query
func (t *Tibber) subscribe(client *graphql.SubscriptionClient, homeID string, done chan error) {
	var query struct {
		tibber.LiveMeasurement `graphql:"liveMeasurement(homeId: $homeId)"`
	}

	var once sync.Once
	// recv := make(chan struct{})
	// errC := make(chan error)

	_, err := client.Subscribe(&query, map[string]any{
		"homeId": graphql.ID(homeID),
	}, func(data []byte, err error) error {
		if err != nil {
			once.Do(func() { done <- err })
		}
		once.Do(func() { close(done) })

		var res struct {
			LiveMeasurement tibber.LiveMeasurement
		}

		if err := json.Unmarshal(data, &res); err != nil {
			t.log.ERROR.Println(err)
			return nil
		}

		// once.Do(func() {
		// 	close(recv)
		// })

		t.mu.Lock()
		t.live = res.LiveMeasurement
		t.updated = time.Now()
		t.mu.Unlock()

		return nil
	})

	if err != nil {
		once.Do(func() { done <- err })
	}

	go func() {
		if err := client.Run(); err != nil {
			once.Do(func() { done <- err })
		}
	}()

	// wait for connection
	// if err == nil {
	// 	select {
	// 	case <-recv:
	// 	case <-time.After(timeout):
	// 		err = api.ErrTimeout
	// 	case err = <-errC:
	// 	}
	// }

	// return err
}

// CurrentPower implements the api.Meter interface
func (t *Tibber) CurrentPower() (float64, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if time.Since(t.updated) > timeout {
		return 0, api.ErrTimeout
	}

	return t.live.Power - t.live.PowerProduction, nil
}

var _ api.PhaseCurrents = (*Tibber)(nil)

// Currents implements the api.PhaseCurrents interface
func (t *Tibber) Currents() (float64, float64, float64, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if time.Since(t.updated) > timeout {
		return 0, 0, 0, api.ErrTimeout
	}

	return t.live.CurrentL1, t.live.CurrentL2, t.live.CurrentL3, nil
}
