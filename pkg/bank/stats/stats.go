package stats

import "github.com/KadamovSh/bank/v2/pkg/bank/types"

// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
    var filtered []types.Payment
    for _, payment := range payments {
        if payment.Category == category {
            filtered = append(filtered, payment)
        }
    }
    return filtered
}