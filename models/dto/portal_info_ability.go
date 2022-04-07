package dto

// /api/ability 科技能力列表数据
type AbilityInfo struct {
	AbilityId        int    `json:"ability_id"`
	AbilityName      string `json:"ability_name"`
	AbilityDesc      string `json:"ability_desc"`
	AbilityIcon      string `json:"ability_icon"`
	AbilityIconHover string `json:"ability_icon_hover"`
}

// /api/skill/detail 科技能力详情数据
type SkillDetail struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	Introduce string `json:"introduce"`
	Website   string `json:"website"`
	Contacts  string `json:"contacts"`
}

// /api/ability/detail 科技能力详情数据
type AbilityDetail struct {
	AbilityId        int            `json:"ability_id"`
	AbilityName      string         `json:"ability_name"`
	AbilityDesc      string         `json:"ability_desc"`
	AbilityIcon      string         `json:"ability_icon"`
	AbilityIconHover string         `json:"ability_icon_hover"`
	AbilityIntro     string         `json:"ability_intro"`
	AbilityWebsite   string         `json:"ability_website"`
	AbilityContacts  string         `json:"ability_contacts"`
	RelatedCase      []*RelatedCase `json:"related_case"`
}

// RelatedCase 关联案例
type RelatedCase struct {
	CaseID int    `json:"case_id"`
	ImgUrl string `json:"img_url"`
	Desc   string `json:"desc"`
	Title  string `json:"title"`
}
