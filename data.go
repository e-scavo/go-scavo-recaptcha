package recaptcha

import "net/http"

/* General Structures */
type Response struct {
	Error_Code     int         `json:"Error_Code" mapstructure:"Error_Code"`
	Error_Dsc      string      `json:"Error_Dsc" mapstructure:"Error_Dsc"`
	Error_Code_Raw int         `json:"Error_Code_Raw" mapstructure:"Error_Code_Raw"`
	Error_Dsc_Raw  string      `json:"Error_Dsc_Raw" mapstructure:"Error_Dsc_Raw"`
	Data           interface{} `json:"data" mapstructure:"Data"`
}

// Recaptcha holds all the necessary options to configure and verify the request.
type Recaptcha struct {
	secret  string
	client  *http.Client
	version int
	action  string
	score   float64
}
