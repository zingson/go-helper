package bankcmbc

import "testing"

func TestRsaVery(t *testing.T) {
	pubKey := `
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAIXj6+vCRpqonhBkGy2c60eZcG7vArM4n5tMMa1f1fJ4bIomQ5bVVoNE
8cdQBc5XTvQx3Tbpk1mr+5jIB1yZXtDp3hIPfAvtkO8m2J8OpQs0Tr5IlNjZdInK
AGGu85zaFXxGvHoW74mtL0JKFPl2Ds2a5ztjK59iacLibIduORb9AgMBAAE=
-----END RSA PUBLIC KEY-----
`
	sign := "79fb217ec2de9cf635c2e0c53b3c029545a1be5fe9dbc9edd681ebe0de571e8746c8ce4533a085dc8a57f0e1d883f24086339ae7f3a0ca1c08ec550587e31954f15c5696482db02f6bd6bbd3ce9fdad69d6bd3559e70ffc1f0c5c1025a1f9e693884edf9b5978c4360b3028308f0ddd3ed5d65a139a58997398f541ffc3ce634"
	err := RsaVery("123", sign, pubKey)
	t.Error(err)
}
