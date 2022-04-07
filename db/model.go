package db

const (
	Accurate       = "accurate" // 精确匹配
	Fuzzy          = "fuzzy"    // 模糊匹配
	Not            = "not"      // 不匹配
	Bwt            = "between"  // 位于两者之间（含）
	In             = "in"       // 在某个范围内
	NotIn          = "notin"    // 不在某个范围内
	Larger         = "larger"   // 大于
	LargerOrEqual  = "le"       // 大于或等于
	Smaller        = "smaller"  // 小于
	SmallerOrEqual = "se"       // 小于或等于
)

// SqlCondition Sql条件
type SqlCondition struct {
	Sql    string   `json:"sql"`
	Values []string `json:"values"` // SQL条件对应的value集合
}

// FieldCondition 字段条件
type FieldCondition struct {
	TableName string `json:"table_name"`
	Field     string `json:"field"`
	Value     string `json:"value"`
	MatchType string `json:"match_type"`
}

// PageCondition 基于Field查询专用的结构体
type PageCondition struct {
	Conditions []FieldCondition `json:"conditions"`
	PageNo     int              `json:"page_no"`
	PageSize   int              `json:"page_size"`
}

// SearchResult 查询返回结果类
type SearchResult struct {
	TotalCount int64                    `json:"total_count"`
	Items      []map[string]interface{} `json:"items"`
}

// SearchModelResult 查询返回结果类
type SearchModelResult struct {
	TotalCount int64       `json:"total_count"`
	Items      interface{} `json:"items"`
}
