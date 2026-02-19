package stats

import "github.com/KadamovSh/bank/v2/pkg/bank/types"

// CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
    totals := make(map[types.Category]types.Money)
    
    for _, payment := range payments {
        totals[payment.Category] += payment.Amount
    }
    
    return totals
}