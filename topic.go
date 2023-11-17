package gokaf

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
)

type topic struct {
	ctx     context.Context
	cancel  context.CancelFunc
	logger  *slog.Logger
	wg      sync.WaitGroup // Add a WaitGroup for synchronization
	name    string
	channel *topicChannel
}

// newTopic creates a new instance of topic.
func newTopic(ctx context.Context, logger *slog.Logger, name string, bufferSize int) *topic {
	ctx, cancel := context.WithCancel(ctx)

	channel := newTopicChannel(bufferSize)

	t := topic{ctx, cancel, logger, sync.WaitGroup{}, name, channel}

	t.wg.Add(1) // Increment the WaitGroup counter
	return &t
}

func (t *topic) close() {
	defer t.wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	defer t.channel.Close()
	// Shutdown. Cancel application context will kill all attached tasks.
	t.logger.Warn(fmt.Sprintf("Topic %s closed", t.name))
	t.cancel()
}

func (t *topic) publish(msg interface{}) error {
	select {
	case <-t.ctx.Done():
		errorMsg := fmt.Sprintf("Topic %s is already closed", t.name)
		t.logger.Warn(errorMsg)
		return fmt.Errorf(errorMsg)
	default:
		if t.channel.IsClosed() {
			e := newTopicClosedError(t.name)
			t.logger.Warn(e.Error())
			return e
		}
		t.channel.ch <- msg
		return nil
	}
}
