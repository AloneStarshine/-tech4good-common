package db

import (
	"git.code.oa.com/fip-team/rasse/redis"
	redigo "github.com/gomodule/redigo/redis"
)

// redis setnx方法 实现数据加锁
func SetNx(key string, value string) (bool, error) {
	cache := redis.GetStdCache()
	reply, err := redigo.Int(cache.Do("SETNX", key, value))
	if err != nil {
		return false, err
	}
	if reply == 1 {
		return true, nil
	}
	return false, nil
}

// redis TTL 获取剩余过期时间
func TTL(key string) (int, error) {
	cache := redis.GetStdCache()
	reply, err := redigo.Int(cache.Do("TTL", key))
	if err != nil || reply == -2{
		return -2 , err
	}
	if reply >= 0 {
		return reply, nil
	}
	return -1 , nil
}

// redis GetString 方法 去掉数据为空判断
func GetString(key string) (string, error) {
	result, err := redis.GetString(key)
	if err == nil {
		return result, nil
	}
	if err.Error() != "redigo: nil returned" {
		return "", err
	} else {
		return result, nil
	}
}

// redis HSet 方法 集成redis hset方法
func HSet (key string , field string , value string) error {
	cache := redis.GetStdCache()
	_ , err := redigo.Int(cache.Do("HSET", key, field ,value))
	if err != nil {
		return err
	}
	return nil
}

// redis HGETALL 方法 集成redis HGETALL 方法
func HGetAll(key string)(map[string]string, error){
	cache := redis.GetStdCache()
	m , err := redigo.StringMap(cache.Do("HGETALL", key))
	if err != nil {
		return nil,err
	}
	return m,nil
}

// redis HGet 方法 集成redis HGet 方法
func HGet(key string,field string)(string, error){
	cache := redis.GetStdCache()
	s , err := redigo.String(cache.Do("HGET", key,field))
	if err != nil {
		return "",err
	}
	return s,nil
}
