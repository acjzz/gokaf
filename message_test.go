package gokaf

import (
	"reflect"
	"testing"
)

func Test_newInternalMessage(t *testing.T) {
	message := "test"
	t.Run("InternalMessage", func(t *testing.T) {
		if got := newInternalMessage(message); !reflect.DeepEqual(got, internalMessage{message}) {
			t.Errorf("newInternalMessage() = %v, want %v", got, internalMessage{message})
		}
	})
}
