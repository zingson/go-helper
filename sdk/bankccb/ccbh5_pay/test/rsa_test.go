package test

import (
	_ "embed"
)

//go:embed .secret/config.json
var cstr string
