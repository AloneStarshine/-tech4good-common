package dto

// /api/home/info 返回数据
type InfoResult struct {
	Introduce *Introduce `json:"introduce"`
	Cases     []*Case    `json:"cases"`
	Abilities []*Ability `json:"abilities"`
}

// /api/home/info 返回数据-role数据
type Role struct {
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	RoleIcon string `json:"role_icon"`
	RoleId   int    `json:"role_id"`
}

// /api/home/info 返回数据-首页介绍数据
type Introduce struct {
	IntroduceDesc string  `json:"introduce_desc"`
	Roles         []*Role `json:"roles"`
}

// /api/home/info 返回数据-首页案例数据
type Case struct {
	CaseLabel string `json:"case_label"`
	CaseTitle string `json:"case_title"`
	CaseDesc  string `json:"case_desc"`
	CaseUrl   string `json:"case_url"`
	CaseId    int    `json:"case_id"`
	CasePath  string `json:"case_path"`
}

// /api/home/info 返回数据-首页科技能力数据
type Ability struct {
	AbilityId        int    `json:"ability_id"`
	AbilityName      string `json:"ability_name"`
	AbilityDesc      string `json:"ability_desc"`
	AbilityIcon      string `json:"ability_icon"`
	AbilityIconHover string `json:"ability_icon_hover"`
}
