package domain

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type PageParams struct {
	Page   int
	Limit  int
	Action string
}

func NormalizePageParams(page, limit int) PageParams {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	return PageParams{Page: page, Limit: limit}
}

func Offset(p PageParams) int {
	return (p.Page - 1) * p.Limit
}
