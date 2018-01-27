package rbac

import (
	"github.com/artpar/rbac/cache"
	"github.com/artpar/rbac/db"
)

type RBACConfig struct {
	Redis *cache.RedisConfig
	Mgo   *db.MgoConf
}
