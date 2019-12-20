package src

import (
	"context"
	"reflect"
	"testing"
)

func Test_newConsumer(t *testing.T) {
	type args struct {
		ctx context.Context
		ch  *chan messageInterface
	}
	tests := []struct {
		name string
		args args
		want *consumer
	}{
		{ "constructor", args{}, nil },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan messageInterface)

			tt.args = args{ ctx, &ch, }
			tt.want = &consumer{ ctx, &ch, }

			if got := newConsumer(tt.args.ctx, tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consumer_Run(t *testing.T) {
	type fields struct {
		ctx     context.Context
		channel *chan messageInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{ "constructor", fields{}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())

			ch := make(chan messageInterface)

			tt.fields = fields{ ctx, &ch, }

			c := &consumer{
				ctx:     tt.fields.ctx,
				channel: tt.fields.channel,
			}

			go func(){
				c.run()
			}()
			cancel()

		})
	}
}
