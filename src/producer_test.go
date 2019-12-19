package src

import (
	"context"
	"reflect"
	"testing"
)

func Test_newProducer(t *testing.T) {
	type args struct {
		ctx     context.Context
		channel *chan messageInterface
	}
	tests := []struct {
		name string
		args args
		want *producer
	}{
		{ "constructor", args{}, nil },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan messageInterface)

			tt.args = args{ ctx, &ch, }
			tt.want = &producer{ ctx, &ch, }

			if got := newProducer(tt.args.ctx, tt.args.channel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_producer_Publish(t *testing.T) {
	type fields struct {
		ctx     context.Context
		channel *chan messageInterface
	}
	type args struct {
		message messageInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{ "publish", fields{}, args{ NewInternalMessage("message10")}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())

			ch := make(chan messageInterface)

			tt.fields = fields{ ctx, &ch, }

			p := &producer{
				ctx:     tt.fields.ctx,
				channel: tt.fields.channel,
			}

			go func(){
				p.Publish(tt.args.message)
			}()
			cancel()
		})
	}
}
