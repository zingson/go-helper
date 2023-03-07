package sha

import "testing"

func TestSha(t *testing.T) {
	t.Log(Sha1("root"))
	t.Log(Sha256("1"))
	t.Log(Sha512("1"))
}
