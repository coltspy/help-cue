package models

import "time"

// HelpRequest struct represents a single student help request
type HelpRequest struct {
	ID          int       `json:"id"`
	StudentName string    `json:"studentName"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "pending", "claimed", "resolved"
	CreatedAt   time.Time `json:"createdAt"`
}
