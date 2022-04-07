package utils

import (
	"git.code.oa.com/fip-team/rasse/orm"
	"git.code.oa.com/fip-team/rasse/redis"
)

// GetString 从redis或db获取数据
func GetString(key string, funcName func(map[string]interface{}) (string, error),
	param map[string]interface{}) (reply string, err error) {

	reply, err = redis.GetString(key)
	if err != nil && err.Error() != "redigo: nil returned" {
		return "", err
	}
	if reply == "" {
		reply, err = funcName(param)
		if err != nil {
			return "", err
		}
		err = redis.SetString(key, reply, 60*24)
		if err != nil {
			return "", err
		}
		return reply, nil
	}
	return reply, nil
}

// Handler 数据增删改结构体
type Handler struct {
	FuncName func(*orm.FiTX, map[string]interface{}, interface{}) error
	Where    map[string]interface{}
	Data     interface{}
}

// MultiHandle 多操作数据库处理方法
func MultiHandle(key string, f []*Handler) error {
	exist, err := redis.Exists(key)
	if err != nil {
		return err
	}
	if exist {
		err = redis.Delete(key)
		if err != nil {
			return err
		}
	}
	tx := orm.BeginTransaction()
	defer tx.EndTransaction()
	for _, item := range f {
		err = item.FuncName(tx, item.Where, item.Data)
		if err != nil {
			tx.RollbackTranction()
			return err
		}
	}
	tx.Commit()
	return nil
}

// Handle 单一操作数据库处理方法
func Handle(key string, funcName func(interface{}) error, data interface{}) error {
	exist, err := redis.Exists(key)
	if err != nil {
		return err
	}
	if exist {
		err = redis.Delete(key)
		if err != nil {
			return err
		}
	}
	err = funcName(data)
	return err
}
