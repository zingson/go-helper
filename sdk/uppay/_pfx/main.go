package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/pkcs12"
	"os"
)

// 解析pfx证书获取pem格式私钥
func main() {
	pfxFile := "D:\\netdisk\\company\\中国银联\\11-银联全渠道支付商户\\银联支付-全渠道-89833027922F01E-广电广通（地铁购票）\\89833027922F01E.pfx" // os.Args[1]  // pfx路径
	password := "123456"                                                                                            //os.Args[2] // pfx密码

	fbytes, err := os.ReadFile(pfxFile)
	if err != nil {
		panic(err)
	}
	priKey, cert, err := pkcs12DecodeAll(fbytes, password)
	if err != nil {
		panic(err)
	}

	// 格式化私钥
	derStream, err := x509.MarshalPKCS8PrivateKey(priKey[0])
	if err != nil {
		panic(err)
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	})

	// 写入到文件
	os.WriteFile("private.key", pemBytes, os.ModeType)

	fmt.Println("*************** 解析银联pfx证书私钥与序列号 ***************")
	fmt.Println("SerialNumber=" + cert[0].SerialNumber.String())
	fmt.Println("私钥证书 \n" + string(pemBytes))
}

// pkcs12DecodeAll extracts all certificate and private keys from pfxData.
func pkcs12DecodeAll(pfxData []byte, password string) ([]interface{}, []*x509.Certificate, error) {
	var privateKeys []interface{}
	var certificates []*x509.Certificate

	blocks, err := pkcs12.ToPEM(pfxData, password)
	if err != nil {
		log.Printf("error while converting to PEM: %s", err)
		return nil, nil, err
	}

	for _, b := range blocks {
		if b.Type == "CERTIFICATE" {
			certs, err := x509.ParseCertificates(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			certificates = append(certificates, certs...)

		} else if b.Type == "PRIVATE KEY" {
			privateKey, err := x509.ParsePKCS1PrivateKey(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}
	return privateKeys, certificates, err
}
