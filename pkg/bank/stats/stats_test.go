package stats

import (
    "testing"
    "github.com/KadamovSh/bank/v2/pkg/bank/types"
	"reflect"
)


func TestFilterByCategory_empty(t *testing.T) {
    // Arrange
    payments := []types.Payment{}
    category := types.Category("auto")

    // Act
    result := FilterByCategory(payments, category)

    // Assert
    if len(result) != 0 {
        t.Errorf("len(result) = %d, want %d", len(result), 0)
    }
}

func TestFilterByCategory_nil(t *testing.T) {
    // Arrange
    var payments []types.Payment  // nil!
    category := types.Category("auto")

    // Act
    result := FilterByCategory(payments, category)

    // Assert
    if len(result) != 0 {
        t.Errorf("len(result) = %d, want %d", len(result), 0)
    }
}

func TestFilterByCategory_notFound(t *testing.T) {
    // Arrange
    payments := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "food", Amount: 2_000_00},
        {ID: 3, Category: "fun", Amount: 3_000_00},
    }
    category := types.Category("mobile")

    // Act
    result := FilterByCategory(payments, category)

    // Assert
    if len(result) != 0 {
        t.Errorf("len(result) = %d, want %d", len(result), 0)
    }
}



func TestFilterByCategory_foundOne(t *testing.T) {
    // Arrange
    payments := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "food", Amount: 2_000_00},
        {ID: 3, Category: "fun", Amount: 3_000_00},
    }
    category := types.Category("food")
    
    expected := []types.Payment{
        {ID: 2, Category: "food", Amount: 2_000_00},
    }

    // Act
    result := FilterByCategory(payments, category)

    // Assert
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("result = %v, want %v", result, expected)
    }
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
    // Arrange
    payments := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "auto", Amount: 2_000_00},
        {ID: 3, Category: "food", Amount: 3_000_00},
        {ID: 4, Category: "auto", Amount: 4_000_00},
        {ID: 5, Category: "fun", Amount: 5_000_00},
    }
    category := types.Category("auto")
    
    expected := []types.Payment{
        {ID: 1, Category: "auto", Amount: 1_000_00},
        {ID: 2, Category: "auto", Amount: 2_000_00},
        {ID: 4, Category: "auto", Amount: 4_000_00},
    }

    // Act
    result := FilterByCategory(payments, category)

    // Assert
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("result = %v, want %v", result, expected)
    }
}