package models

type Host struct {
	Host string `json:"host"`
	Port int16 `json:"port"`
	Protocol string `json:"protocol"`
	IsPublic bool `json:"isPublic"`
	Status string `json:"status"`
	StartTime int64 `json:"startTime"`
	TestTime int64 `json:"testTime"`
	EngineVersion string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints []Endpoint `json:"endpoints"`
}