/*
 * @Author: huqing
 * @Date: 2019-02-15 14:21:22
 * @Last Modified by: huqing
 * @Last Modified time: 2019-02-15 14:24:30
 * @desc:
 */

package myRedis

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

/// 活检
func PingSvr(clt *redis.Client) error {
	defer func() {
		recover()
	}()

	if clt == nil {
		return fmt.Errorf("hand for redis connector is empty")
	}

	res, err := clt.Ping().Result()
	if err != nil || strings.Compare(res, "PONG") != 0 {
		return err
	}
	return nil
}

/// 连接
func ConnectRds(opt *redis.Options) (*redis.Client, error) {
	if opt == nil {
		return nil, fmt.Errorf("redis's config is nil")
	}

	clt := redis.NewClient(opt)

	if err := PingSvr(clt); err != nil {
		return nil, err
	}
	return clt, nil
}

/// 关闭连接
func CloseRds(clt *redis.Client) {
	if clt != nil {
		clt.Close()
		clt = nil
	}
}
