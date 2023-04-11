package models

type Coffee struct {
	Name            string `json:"name"`
	Brand           string `json:"brand"`
	Recommendations int32  `json:"recommendations"`
}

type QueryCoffees struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
