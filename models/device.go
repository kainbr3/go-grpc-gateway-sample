package models

//Device represents a smart home device
type Device struct {
	ID       int    `json:"id"`
	Hardware string `json:"hardware"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	State    int    `json:"state"`
}
