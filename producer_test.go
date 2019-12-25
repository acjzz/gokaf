package gokaf

import (
"context"
	"fmt"
	"strings"
	"testing"
)

func Test_producer(t *testing.T) {
	t.Run("Producer", func(t *testing.T) {
		wantErrMsg := "topic closed"
		ctx, cancel := context.WithCancel(context.Background())
		var channel chan internalMessage
		p := newProducer(ctx, &channel)
		cancel()
		err :=p.publish(internalMessage{"test"})
		if err == nil {
			t.Errorf("publish() wanted error")
		} else {
			ErrMsg := fmt.Sprintf("%v", err)
			if strings.Compare(ErrMsg, wantErrMsg) != 0 {
				t.Errorf("publish() error = '%v', wantErr '%v'", ErrMsg, wantErrMsg)
			}
		}
	})
}
