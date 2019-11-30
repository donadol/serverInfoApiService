package models

type Domain struct {
	Host string `json:"host"`
	ServersChanged bool `json:"servers_changed"`
	MinGrade string `json:"ssl_grade"`
	PreviousMinGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown string `json:"is_down"`
	Servers []Server `json:"servers"`
}
