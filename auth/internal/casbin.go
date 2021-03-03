package internal

import (
	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
)

type CastbinConf struct {
	ModelPath    string
	PolicyDriver string
	PolicyPath   string
	DbConf       DatabaseConf
}

var gEnforcer *casbin.Enforcer

func InitCasbin(conf CastbinConf) {
	var err error
	switch conf.PolicyDriver {
	case "", "file":
		gEnforcer, err = casbin.NewEnforcer(conf.ModelPath, conf.PolicyPath)
		if err != nil {
			panic(err)
		}
	case "db":
		db, err := NewDb(conf.DbConf)
		if err != nil {
			panic(err)
		}
		a, err := sqladapter.NewAdapter(db, conf.DbConf.Driver, "casbin_rule")
		if err != nil {
			panic(err)
		}
		gEnforcer, err = casbin.NewEnforcer(conf.ModelPath, a)
		if err != nil {
			panic(err)
		}
	}
}

func CheckAccess(sub, obj, act string) bool {
	if gEnforcer == nil {
		panic("castbin undefined")
	}
	ok, err := gEnforcer.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}
	return ok
}
