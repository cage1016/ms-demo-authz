package postgres

import (
	"context"
	"time"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus" // TODO change log system with gokit log

	storageV1 "github.com/cage1016/gokit-istio-security/internal/app/authz/storage/v1"
)

type policyChangeNotifier struct {
	connInfo             string
	minReconnectInterval time.Duration
	maxReconnectInterval time.Duration
	pingInterval         time.Duration
	notificationChan     chan storageV1.PolicyChangeNotification
	shutdown             func()
}

func newPolicyChangeNotifier(ctx context.Context, connInfo string) (storageV1.PolicyChangeNotifier, error) {
	ctx, cancel := context.WithCancel(ctx)
	p := &policyChangeNotifier{
		connInfo:             connInfo,
		minReconnectInterval: 10 * time.Second,
		maxReconnectInterval: time.Minute,
		pingInterval:         10 * time.Second,
		notificationChan:     make(chan storageV1.PolicyChangeNotification, 1),
		shutdown:             cancel,
	}
	listener := pq.NewListener(p.connInfo, p.minReconnectInterval, p.maxReconnectInterval, nil)
	err := listener.Listen("policychange")
	if err != nil {
		return nil, err
	}

	go p.run(ctx, listener)
	return p, nil
}

func (p *policyChangeNotifier) C() <-chan storageV1.PolicyChangeNotification {
	return p.notificationChan
}

func (p *policyChangeNotifier) Close() error {
	p.shutdown()
	return nil
}

func (p *policyChangeNotifier) run(ctx context.Context, listener *pq.Listener) {

RUNLOOP:
	for {
		select {
		case <-ctx.Done():
			break RUNLOOP
		case n := <-listener.Notify:
			if n == nil {
				continue
			}
			select {
			case p.notificationChan <- storageV1.PolicyChangeNotification{}:
				logrus.Info("Accepted notification from postgres")
			default:
				logrus.Debug("Notification listener mailbox full")
			}
		case <-time.After(p.pingInterval):
			err := listener.Ping()
			if err != nil {
				logrus.WithError(err).Warn("Notification listener failed to ping database")
			}
		}
	}
	if err := listener.Close(); err != nil {
		logrus.WithError(err).Warn("Failed to close notification listener")
	}
}
