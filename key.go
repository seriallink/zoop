package zoop

import "time"

type Key struct {
	Id                       string      `json:"id"`
	Type                     string      `json:"type"`
	Name                     string      `json:"name"`
	TestKey                  string      `json:"test_key"`
	PublishableTestKey       string      `json:"publishable_test_key"`
	ProductionKey            string      `json:"production_key"`
	PublishableProductionKey string      `json:"publishable_production_key"`
	Metadata                 interface{} `json:"metadata"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Info
}
