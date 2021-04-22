package service

import (
	"github.com/go-redis/redis"
	"github/qm012/nacos-adress/global"
)

type cluster struct {
	rdb      *redis.Client
	cacheKey string
}

var (
	_ Storage = &cluster{}
)

func newCluster() Storage {
	c := &cluster{
		rdb:      global.Rdb,
		cacheKey: "nacos:clusterIps",
	}
	return c
}

func (c *cluster) get() ([]string, error) {
	return c.rdb.SMembers(c.cacheKey).Result()
}

func (c *cluster) add(strings []string) error {
	_, err := c.rdb.SAdd(c.cacheKey, strings).Result()
	return err
}

func (c *cluster) delete(strings []string) error {
	_, err := c.rdb.SRem(c.cacheKey, strings).Result()
	return err
}

func (c *cluster) deleteAll() error {
	_, err := c.rdb.Del(c.cacheKey).Result()
	return err
}
