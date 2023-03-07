package bankcmbc

import (
	"crypto/rand"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"testing"
)

func TestGenKey(t *testing.T) {
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	priKeyHex := x509.WritePrivateKeyToHex(privateKey)
	pubKeyHex := x509.WritePublicKeyToHex(&privateKey.PublicKey)
	fmt.Println("私钥：", priKeyHex)
	fmt.Println("公钥：", pubKeyHex)
	/*
		私钥： 1910a235a2bd1c6b8fdf9c2503f91a2474aa00cc3507545515f4d3fd2a43288b
		公钥： 04e47cacd84ba407675beaafc5b150a7985de4d387379decab60dcc91b680929d173c7a267d61802c67020d1500dc69e32056ae5f96e09b4b905dca0bb425290ac
	*/
}
