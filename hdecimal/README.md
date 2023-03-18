# decimal




## Usage

```go
package main

import (
	"fmt"
	"github.com/zingson/go-helper/hdecimal"
)

func main() {
	price, err := hdecimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := hdecimal.NewFromInt(3)

	fee, _ := hdecimal.NewFromString(".035")
	taxRate, _ := hdecimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(hdecimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(hdecimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
```



# 引用项目

- `github.com/shopspring/decimal` 
