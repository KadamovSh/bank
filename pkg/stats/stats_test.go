package stats

import (
    "reflect"
    "testing"
    "github.com/KadamovSh/bank/v2/pkg/bank/types"
)

func TestCategoriesTotal(t *testing.T) {
    payments := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "food", Amount: 2_000_00},
        {ID: 3, Category: "auto", Amount: 3_000_00},
    }
    
    expected := map[types.Category]types.Money{
        "auto": 4_000_00,
        "food": 2_000_00,
    }
    
    result := CategoriesTotal(payments)
    
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("получили %v, ожидали %v", result, expected)
    }
}