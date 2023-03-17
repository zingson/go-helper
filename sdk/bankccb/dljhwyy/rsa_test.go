package dljhwyy

import "testing"

func TestDecrypt(t *testing.T) {
	prikey := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDVbauPnN4KB2YW
hyRymnb/csBJWD/aZyuaalwG6DiSasSF7nxDOdT2ghQsiiH5TvX2Yp10tVR3LI/h
5uwsYYt90FchlIx5jX7ZuU77cxS+dKc5StPOQgIGNe8gO84fByNmlX3h2rzSbjNV
9njeYdbTpyOTMde2FB8T6mJM0BbrX72EkXMl6Kc1P1C88kEQTrTxLFJRxwSboT1Y
jfB7bQc8TOODnjwcYqOcIz00cEhUiWi8/DETqxzHzC3RUbXg7cslqqTSON+StGH2
ERWkIHAbX7afPS373AERCb9UJftUv7nrs7KVYTJEwA7aqR2HhIhobw5oLJ5r9Dnj
YijxTdFrAgMBAAECggEABy11A5Nm9Ddje4Z391Kyhcy6Ir1RCGtH0B2bkq/klyf4
C/kFPM2JF/Ev9H+AvP2mz+5pFS+z834QKKy3bJarNkP3ai2wu7XCelf9C//GxtDt
fsPBc8JMhyDxNchNGkYHLsAAR8QvbXQ/TbjIP9JSgzOvwpd+haUPln/fZm3pF2lF
H0+89jYVhOGDkZh/7hHb6B6/gqxwm5bktgOyvKMP4zPS00WdTDuKljzkHvD9SHpl
2+7BXI0EyrkaoFx5g0Gosx2yMQQwMRyl/zoE9xt2+gBJoOHYIKF/ARRdUcMbZsL/
GhPk5QsgrNMjV8/EtsCle87A1bMCB/4kKkWORDcTgQKBgQDxorG/Fr6h6OoM+eb9
5DrzEpK7e+EcrF0ct+fi5h/5uzyOGvl0VzPS6uYTbbbqF4VPX01eO6jMRc20TF2S
/FMgI73DuKEaqf/rbUxdY7JuoTlKKZp6l25Yfcn8Ef0UZvZx/YqrgsGvr18FbUA5
vJbItLCKSPR7Aw+d27cNK1h7+wKBgQDiHbVkYnCsNpvVScXbuvgdoeRvLbcQWi8v
Y7y7Tu+a8nZP8EI3htPjiODlf4KEoWil3zzYCPswes3/nxODGTOHm27Rc+pUUonS
ip/+RuJheyou/sG2W2mE9zaoEGWhfCSXt+rjp0uuQyoxgX28eP/XWDBFWNljpFfQ
e573U/kVUQKBgDfkMQcJRMkArfqYZ1S80OHenvAc5GIUgGkfJNT389d49rXMCX1e
mk8kJFnEG/YjwKMuYR/13sJ5Nu6IZ3CTSsr5WXCHKUSYo+iZ3HsRj3lFc5Hy85d9
pXl0hx72Ljrk4gZV5hhr40822lzP+3YfaCU+nkMuQtadS6Rw6PNz3kC3AoGAJXge
0t+tJBx4fkOXUe4Np3toSzQcHc2T/Bpe7/sIoXiEOoLptiHVguLvwZf7nNbSbIot
nu7+EO6IrE9EAHlwnIwZNQQsVITI6eam1JASe2zZdKgqmXlUZwBAQmFVNglIVwvX
FJpoZBwlJcb7evviCWFHvnYWr/hPxPB3SyzTHcECgYEAwyKmG1rL0K0EZZgfKIHd
K3HYF9z+ISohpPuND1BSEUh/uJe5h0LYDQeWC3ipzoWSMtAkX6upEk1RF3bQp0PW
o/C/dsTwr8F4Da65Ceyfjs45SmvLOT+GSaH/z66WyCf4ptCUcpFoNJnOn1hEjZfj
zivYMzto5Tg+iXDVGgq0RrQ=
-----END PRIVATE KEY-----
`
	ciphertext := "TEmPylaPf3+CkReSsKVsfx97cKp3zSpylGVjmIWASvbKAVLij3TIkU9ua06etKJ93Sp7jZH3rLh5FT+UtSLSEKpnL0cbZhS41b0V+vnstrnC67QmZd4Hd9DR+0UVzKJ7pv/0SqKWjqCbGB/NEakttAdWOvA4IbWQb4lQAF5wLcUFMdjzT4zezGFGb4gz4DmlfsyvHUV0McqgVqrMhWpvmUBrKZr8/q+IqKxXaeUU9GpqmkhWvfIscQmOoW1M0uhI7F99pcG4VFgt3WY/zK2uB6W9lg6j5fKMLs5lMHD+B2ZFnla6mTwT7H5gICXQGy+qaFIQom82SpyLB5JPcBEhkA=="
	v, err := RsaDecode(ciphertext, prikey)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(v)
}

func TestEncode(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1W2rj5zeCgdmFockcpp2\n/3LASVg/2mcrmmpcBug4kmrEhe58QznU9oIULIoh+U719mKddLVUdyyP4ebsLGGL\nfdBXIZSMeY1+2blO+3MUvnSnOUrTzkICBjXvIDvOHwcjZpV94dq80m4zVfZ43mHW\n06cjkzHXthQfE+piTNAW61+9hJFzJeinNT9QvPJBEE608SxSUccEm6E9WI3we20H\nPEzjg548HGKjnCM9NHBIVIlovPwxE6scx8wt0VG14O3LJaqk0jjfkrRh9hEVpCBw\nG1+2nz0t+9wBEQm/VCX7VL+567OylWEyRMAO2qkdh4SIaG8OaCyea/Q542Io8U3R\nawIDAQAB\n-----END PUBLIC KEY-----\n"
	v, err := RsaEncode(`{"user_id":"testuser"}`, pubKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(v)
}
