package src

import (
	"reflect"
	"testing"
)

func TestNewInternalMessage(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *internalMessage
	}{
		{ "constructor", args{ "message00", }, &internalMessage{"message00"}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInternalMessage(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInternalMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalMessage_consume(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{ "consume", fields{"message10" } },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &internalMessage{
				value: tt.fields.value,
			}
			m.consume()
		})
	}
}
