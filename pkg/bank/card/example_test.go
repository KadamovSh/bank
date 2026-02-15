package card

import (
    "fmt"
    "github.com/KadamovSh/bank/pkg/bank/types"
)

func ExampleTotal() {
    cards := []types.Card{
        {
            Balance: 1_000_00,
            Active:  true,
        },
    }
    fmt.Println(Total(cards))
    // Output: 100000
}