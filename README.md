# A Zoop SDK In Golang #

## Get It ##

Use `go get -u github.com/seriallink/zoop` to get or update it.

## Usage ##

### Quick start ###

```go
package main

import (
    "fmt"
    "github.com/seriallink/zoop"
)

func main() {
    
    client = zoop.NewClient("<ApiKey>", "<MarketPlaceId>", "<SellerId>")
    
    params := &zoop.CreditCardParams{
        HolderName: "John Doe",
        ExpirationMonth: "10",
        ExpirationYear: "20",
        CardNumber: "4539003370725497",
        SecurityCode: "123",
    }

    token, err := client.NewCardToken(params)
    if err != nil {
        panic(err)
    }
  
    fmt.Println("Token ID:", token.Id)
    fmt.Println("Card ID:", token.Card.Id)
  
}
```
