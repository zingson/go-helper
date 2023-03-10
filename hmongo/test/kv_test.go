package test

import (
	_ "embed"
	"github.com/zingson/go-helper/hmongo"
	"testing"
)

func TestNew(t *testing.T) {

	t.Log("........")
	c := hmongo.NewKV(hmongo.DsnLoad(), hmongo.KvName(), "mongo.himkt").Watch()
	t.Log(c.GetCache())

}
