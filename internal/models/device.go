package models

type Device struct {
	Name       string `json:"name"`
	IpAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
	Type       string `json:"type"`
	Connected  bool   `json:"connected"`
}
