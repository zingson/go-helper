package example

import (
	"github.com/zingson/go-helper/sdk/uppay/ocapp"
	"testing"
)

func TestRsaWithSha256Sign(t *testing.T) {
	sign, err := ocapp.RsaWithSha256Sign("1", cfg89833027372F284.MerPrivateKey)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(sign)
}

//btzAR8D+2uEnkTBWd6F9j0KFACj9H1LonX7lF1d15tpFjJQqbQzx2Kqe7BoMkCsjwxlDrX3MNX5vgbqXwJH2YHU9DXrHgVX16tLpOznIthhzdZbl8kgLA/ocOptDojASUVXBLZYkRIrRWlcum1g++S4/G4wN6b0kpaxA7P5QKvGdr2L3+LDavO/XWrKyP4vHE6PvDZ9O3sXhUnPn86+Hi50PKFsqh6NwSohYTRwUl15t/EPW0CbDAzZ2JJ5lXdXUT72vVhfMN5SieFHS7hiBlPRowxocSjv9heWHWrJHJt4DEowUnQh8oChoLQrLYE2lCnlFGGRFpdOwd3Ozq8O/Uw==
//btzAR8D+2uEnkTBWd6F9j0KFACj9H1LonX7lF1d15tpFjJQqbQzx2Kqe7BoMkCsjwxlDrX3MNX5vgbqXwJH2YHU9DXrHgVX16tLpOznIthhzdZbl8kgLA/ocOptDojASUVXBLZYkRIrRWlcum1g++S4/G4wN6b0kpaxA7P5QKvGdr2L3+LDavO/XWrKyP4vHE6PvDZ9O3sXhUnPn86+Hi50PKFsqh6NwSohYTRwUl15t/EPW0CbDAzZ2JJ5lXdXUT72vVhfMN5SieFHS7hiBlPRowxocSjv9heWHWrJHJt4DEowUnQh8oChoLQrLYE2lCnlFGGRFpdOwd3Ozq8O/Uw==
