package validation

import (
	"github.com/jesusGZ/simple_server_of_fiber/pkg/entity"
)

// CreateOrUpdate to verify the required fields
func CreateOrUpdate(expense entity.Expense) []string {
	var messages []string

	if !isValid(expense.Title) {
		messages = append(messages, "Title is required")
	}

	if !isValid(expense.Total) {
		messages = append(messages, "Total should be greater than zero")
	}

	return messages
}
