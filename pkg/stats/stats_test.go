package stats

import (
    "reflect"
    "testing"
    "github.com/KadamovSh/bank/v2/pkg/bank/types"
)

func TestCategoriesTotal(t *testing.T) {
    // Arrange
    payments := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "food", Amount: 2_000_00},
        {ID: 3, Category: "auto", Amount: 3_000_00},
        {ID: 4, Category: "auto", Amount: 4_000_00},
        {ID: 5, Category: "fun", Amount: 5_000_00},
    }
    
    expected := map[types.Category]types.Money{
        "auto": 8_000_00,
        "food": 2_000_00,
        "fun":  5_000_00,
    }

    // Act
    result := CategoriesTotal(payments)

    // Assert
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("CategoriesTotal() = %v, want %v", result, expected)
    }
}