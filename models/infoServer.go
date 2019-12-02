package models

type InfoServer struct {
	ServersChanged bool `json:"servers_changed"`
	MinGrade string `json:"ssl_grade"`
	PreviousMinGrade string `json:"previous_ssl_grade"`
	Logo string `json:"logo"`
	Title string `json:"title"`
	IsDown bool `json:"is_down"`
	Servers []Server `json:"servers"`
}