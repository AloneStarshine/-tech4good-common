package db

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"git.code.oa.com/tech4good/common/utils"
)

// BuildCondition 根据condition生成where条件
func BuildCondition(scons []SqlCondition, fcons []FieldCondition) (string, []interface{}, error) {
	sql := ""
	pattern := "^[a-z0-9A-Z_\\.]+$"
	var params []interface{}
	// 添加sql条件
	for _, con := range scons {
		sql += " and (" + con.Sql + ")"
		for _, item := range con.Values {
			params = append(params, item)
		}
	}
	// 添加字段条件
	for _, con := range fcons {

		if ok, _ := regexp.MatchString(pattern, con.Field); !ok {
			return "", params, errors.New("列名错误:" + con.Field)
		}
		tableAlias := ""
		if con.TableName != "" {
			tableAlias = con.TableName + "."
		}
		if con.MatchType == Fuzzy {
			sql += fmt.Sprintf(" and upper(%s`%s`) like ? ", tableAlias, con.Field)
			params = append(params, "%"+strings.ToUpper(fmt.Sprintf("%v", con.Value))+"%")
		} else if con.MatchType == Accurate || con.MatchType == "" {
			sql += fmt.Sprintf(" and %s`%s` = ? ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == Not {
			sql += fmt.Sprintf(" and %s`%s` != ? ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == In {
			sql += fmt.Sprintf(" and %s`%s` in (?) ", tableAlias, con.Field)
			params = append(params, utils.Split(con.Value, ";", ","))
		} else if con.MatchType == Larger {
			sql += fmt.Sprintf(" and %s`%s` > (?) ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == LargerOrEqual {
			sql += fmt.Sprintf(" and %s`%s` >= (?) ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == Smaller {
			sql += fmt.Sprintf(" and %s`%s` < (?) ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == SmallerOrEqual {
			sql += fmt.Sprintf(" and %s`%s` <= (?) ", tableAlias, con.Field)
			params = append(params, con.Value)
		} else if con.MatchType == Bwt {
			sql += fmt.Sprintf(" and ( %s`%s` between ? and ? ) ", tableAlias, con.Field)
			condition := fmt.Sprintf("%v", con.Value)
			values := utils.Split(condition, ";", ",")
			var start, end string
			if len(values) >= 2 {
				start = values[0]
				end = values[1]
			} else {
				start = con.Value
				end = ""
			}
			params = append(params, start)
			params = append(params, end)
		}
	}
	return sql, params, nil
}
