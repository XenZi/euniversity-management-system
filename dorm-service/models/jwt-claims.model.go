package models

type Person struct {
	Name  string   `json:"name"`
	PID   string   `json:"pid"`
	Roles []string `json:"roles"`
}
