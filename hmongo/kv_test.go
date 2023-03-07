package hmongo

import (
	_ "embed"
	"testing"
)

func TestNew(t *testing.T) {

	t.Log("........")
	c := NewKV(DsnLoad(), KvName(), "mongo.himkt").Watch()
	t.Log(c.GetCache())

}
