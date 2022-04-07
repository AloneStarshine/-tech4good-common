package db

import (
	"fmt"

	"git.code.oa.com/fip-team/rasse/orm"
)

// Pager 表分页查询的方法
func Pager(tableName string, scons []SqlCondition, fcons []FieldCondition,
	pageSize int64, pageIndex int64, orderBy string, traceID string) (SearchResult, error) {
	var result SearchResult
	sql, params, offset, err := pageSql(tableName, scons, fcons, &pageSize, &pageIndex, &orderBy)
	countQ := orm.NewQuery().WithLogTraceID(traceID)
	countSql := fmt.Sprintf(sql, " count(1) as count ")
	err = countQ.ExecuteTextQuery(&result, countSql, params...)
	dataQ := orm.NewQuery().WithLogTraceID(traceID)
	sql = fmt.Sprintf(sql, " * ") + " order by " + orderBy + fmt.Sprintf(" limit %v offset %v ", pageSize, offset)
	rows, err := dataQ.Raw(sql, params...).Rows()
	defer rows.Close()
	if err != nil {
		return result, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return result, err
	}
	if err != nil {
		return result, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		err = rows.Scan(valuePtrs...)
		if err != nil {
			return result, err
		}
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		result.Items = append(result.Items, entry)
	}
	return result, nil
}

// PagerModel 表分页查询的方法
func PagerModel(tableName string, scons []SqlCondition, fcons []FieldCondition,
	pageSize int64, pageIndex int64, orderBy string, traceID string) (SearchModelResult, error) {
	var result SearchModelResult
	sql, params, offset, err := pageSql(tableName, scons, fcons, &pageSize, &pageIndex, &orderBy)
	countQ := orm.NewQuery().WithLogTraceID(traceID)
	countSql := fmt.Sprintf(sql, " count(1) as count ")
	err = countQ.ExecuteTextQuery(&result, countSql, params...)
	if err != nil {
		return result, err
	}
	dataQ := orm.NewQuery().WithLogTraceID(traceID)
	sql = fmt.Sprintf(sql, " * ") + " order by " + orderBy + fmt.Sprintf(" limit %v offset %v ", pageSize, offset)
	//println("sql：" + sql)
	err = dataQ.ExecuteTextQuery(result.Items, sql, params...)
	if err != nil {
		return result, err
	}
	return result, nil
}

// pageSql 拼接sql语句
func pageSql(tableName string, scons []SqlCondition, fcons []FieldCondition,
	pageSize *int64, pageIndex *int64, orderBy *string) (string, []interface{}, int64, error) {
	if *orderBy == "" {
		*orderBy = " id desc "
	}
	if *pageIndex == 0 {
		*pageIndex = 1
	}
	if *pageSize == 0 {
		*pageSize = 10
	}
	var offset int64
	if *pageIndex == 1 {
		offset = 0
	} else {
		offset = (*pageIndex - 1) * (*pageSize)
	}
	sql := "select %s from " + tableName + " where 1=1 "
	whereSql, params, err := BuildCondition(scons, fcons)
	if err == nil {
		sql = sql + whereSql
	}
	return sql, params, offset, err
}
