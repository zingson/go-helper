package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("k%d", i)

		client.Set(context.Background(), key, key, time.Minute)

		t.Logf(client.Get(context.Background(), key).Val())
	}
}
