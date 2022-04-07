package dto

// /api/instance 案例列表数据
type CaseList struct {
	CategoryId   int         `json:"category_id"`
	CategoryName string      `json:"category_name"`
	CaseList     []*CaseInfo `json:"case_list"`
}

//案例数据
type CaseInfo struct {
	CaseId int    `json:"case_id"`
	ImgUrl string `json:"img_url"`
	Desc   string `json:"desc"`
	Title  string `json:"title"`
}

// /api/case/detail/:id 案例详情
type CaseDetail struct {
	CaseId           int        `json:"case_id"`
	CaseName         string     `json:"case_name"`
	CategoryId       int        `json:"category_id"`
	CategoryName     string     `json:"category_name"`
	CaseImg          string     `json:"case_img"`
	CaseImgDesc      string     `json:"case_img_desc"`
	CaseDesc         string     `json:"case_desc"`
	CaseIntro        string     `json:"case_intro"`
	CaseApplication  string     `json:"case_application"`
	CaseSignificance string     `json:"case_significance"`
	CaseWebsite      string     `json:"case_website"`
	CaseContacts     string     `json:"case_contacts"`
	CaseTags         []*CaseTag `json:"case_tags"`
}

// 案例应用技术能力
type CaseTag struct {
	AbilityId   int    `json:"ability_id"`
	AbilityName string `json:"ability_name"`
	HasURL      bool   `json:"hasURL"`
	Url         string `json:"url"`
}
