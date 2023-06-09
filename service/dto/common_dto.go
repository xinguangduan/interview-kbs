package dto

// ===============================================================================
// = 通用ID对应的DTO
type CommonDTO struct {
	Uid string `json:"uid" form:"uid" uri:"uid"`
}

// ===============================================================================
// = 分页对应的DTO
type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (m *Paginate) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}

	return m.Page
}

func (m *Paginate) GetLimit() int {
	if m.Limit <= 0 {
		m.Limit = 10
	}

	return m.Limit
}
