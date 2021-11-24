package models

import "ottosfa-api-web/database/dbmodels"

// ReqFilterJobCategories ..
type ReqFilterJobCategories struct {
	ID    int64  `json:"id" example:"1"`
	Name  string `json:"name" example:"Kategori A"`
	Limit int    `json:"limit" example:"10"`
	Page  int    `json:"page"  example:"1"`
}

type ResFilterJobCategories struct {
	JobCategories []dbmodels.JobCategories `json:"data"`
	Total         int                      `json:"total"`
}
