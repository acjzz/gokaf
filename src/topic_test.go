package src

import (
	"context"
	"reflect"
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
			tt.want = &Topic{ ctx, "topic", ch, }

			got := NewTopic(tt.args.ctx, tt.args.name);
			// ---------------------------------------------------------------------------------------------------------
			// HackAround to make this test work, as channel is instantiated by the constructor
			// ---------------------------------------------------------------------------------------------------------
			tt.want.channel = got.channel
			// ---------------------------------------------------------------------------------------------------------
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_NewConsumer(t1 *testing.T) {
	type fields struct {
		ctx     context.Context
		name    string
		channel chan messageInterface
	}
	tests := []struct {
		name   string
		fields fields
		want   *consumer
	}{
		{ "consumer from topic", fields{}, &consumer{ nil, nil}, },
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan messageInterface)

			tt.fields.ctx = ctx
			tt.fields.channel = ch

			t := &Topic{
				ctx:     tt.fields.ctx,
				name:    tt.fields.name,
				channel: tt.fields.channel,
			}

			tt.want.ctx = tt.fields.ctx
			tt.want.channel = &tt.fields.channel

			if got := t.NewConsumer(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopic_NewProducer(t1 *testing.T) {
	type fields struct {
		ctx     context.Context
		name    string
		channel chan messageInterface
	}
	tests := []struct {
		name   string
		fields fields
		want   *producer
	}{
		{ "producer from topic", fields{}, &producer{ nil, nil}, },
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ch := make(chan messageInterface)

			tt.fields.ctx = ctx
			tt.fields.channel = ch

			t := &Topic{
				ctx:     tt.fields.ctx,
				name:    tt.fields.name,
				channel: tt.fields.channel,
			}

			tt.want.ctx = tt.fields.ctx
			tt.want.channel = &tt.fields.channel

			if got := t.NewProducer(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("NewProducer() = %v, want %v", got, tt.want)
			}
		})
	}
}
