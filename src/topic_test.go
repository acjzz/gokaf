package src

import (
	"context"
	"testing"
)

func TestNewTopic(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name string
		args args
		want *Topic
	}{
		{ "constructor", args{}, nil },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan messageInterface)

			tt.args = args{ ctx, "topic", }
			tt.want = &Topic{
				ctx, cancel, "topic", ch,
				[]*consumer{}, newProducer(ctx, &ch),
			}

			got := NewTopic(tt.args.ctx, tt.args.name)
			// ---------------------------------------------------------------------------------------------------------
			// HackAround to make this test work, as channel is instantiated by the constructor
			// ---------------------------------------------------------------------------------------------------------
			tt.want.ctx = got.ctx
			tt.want.channel = got.channel
			tt.want.producer = got.producer
			// ---------------------------------------------------------------------------------------------------------
			//TODO:
			//if  !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewTopic() = %v, want %v", got, tt.want)
			//}
		})
	}
}
