// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gdb provides ORM features for popular relationship databases.
package gdb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/snail007/gf/os/glog"

	"github.com/snail007/gf/container/gmap"
	"github.com/snail007/gf/container/gring"
	"github.com/snail007/gf/container/gtype"
	"github.com/snail007/gf/container/gvar"
	"github.com/snail007/gf/os/gcache"
	"github.com/snail007/gf/util/grand"
)

// 数据库操作接口
type DB interface {
	// 建立数据库连接方法(开发者一般不需要直接调用)
	Open(config *ConfigNode) (*sql.DB, error)

	// SQL操作方法 API
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string, execOnMaster ...bool) (*sql.Stmt, error)

	// 内部实现API的方法(不同数据库可覆盖这些方法实现自定义的操作)
	doQuery(link dbLink, query string, args ...interface{}) (rows *sql.Rows, err error)
	doGetAll(link dbLink, query string, args ...interface{}) (result Result, err error)
	doExec(link dbLink, query string, args ...interface{}) (result sql.Result, err error)
	doPrepare(link dbLink, query string) (*sql.Stmt, error)
	doInsert(link dbLink, table string, data interface{}, option int, batch ...int) (result sql.Result, err error)
	doBatchInsert(link dbLink, table string, list interface{}, option int, batch ...int) (result sql.Result, err error)
	doUpdate(link dbLink, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error)
	doDelete(link dbLink, table string, condition string, args ...interface{}) (result sql.Result, err error)

	// 数据库查询
	GetAll(query string, args ...interface{}) (Result, error)
	GetOne(query string, args ...interface{}) (Record, error)
	GetValue(query string, args ...interface{}) (Value, error)
	GetCount(query string, args ...interface{}) (int, error)
	GetStruct(objPointer interface{}, query string, args ...interface{}) error
	GetStructs(objPointerSlice interface{}, query string, args ...interface{}) error
	GetScan(objPointer interface{}, query string, args ...interface{}) error

	// 创建底层数据库master/slave链接对象
	Master() (*sql.DB, error)
	Slave() (*sql.DB, error)

	// Ping
	PingMaster() error
	PingSlave() error

	// 开启事务操作
	Begin() (*TX, error)

	// 数据表插入/更新/保存操作
	Insert(table string, data interface{}, batch ...int) (sql.Result, error)
	Replace(table string, data interface{}, batch ...int) (sql.Result, error)
	Save(table string, data interface{}, batch ...int) (sql.Result, error)

	// 数据表插入/更新/保存操作(批量)
	BatchInsert(table string, list interface{}, batch ...int) (sql.Result, error)
	BatchReplace(table string, list interface{}, batch ...int) (sql.Result, error)
	BatchSave(table string, list interface{}, batch ...int) (sql.Result, error)

	// 数据修改/删除
	Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)

	// 创建链式操作对象
	From(tables string) *Model
	Table(tables string) *Model

	// 设置管理
	SetDebug(debug bool)
	SetSchema(schema string)
	GetQueriedSqls() []*Sql
	GetLastSql() *Sql
	PrintQueriedSqls()
	SetLogger(logger *glog.Logger)
	SetMaxIdleConnCount(n int)
	SetMaxOpenConnCount(n int)
	SetMaxConnLifetime(n int)
	Tables() (tables []string, err error)
	TableFields(table string) (map[string]*TableField, error)

	// 内部方法接口
	getCache() *gcache.Cache
	getChars() (charLeft string, charRight string)
	getDebug() bool
	quoteWord(s string) string
	setSchema(sqlDb *sql.DB, schema string) error
	filterFields(table string, data map[string]interface{}) map[string]interface{}
	formatWhere(where interface{}, args []interface{}) (newWhere string, newArgs []interface{})
	convertValue(fieldValue []byte, fieldType string) interface{}
	rowsToResult(rows *sql.Rows) (Result, error)
	handleSqlBeforeExec(sql string) string
}

// 执行底层数据库操作的核心接口
type dbLink interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*sql.Stmt, error)
}

// 数据库链接对象
type dbBase struct {
	db               DB                           // 数据库对象
	group            string                       // 配置分组名称
	debug            *gtype.Bool                  // (默认关闭)是否开启调试模式，当开启时会启用一些调试特性
	sqls             *gring.Ring                  // (debug=true时有效)已执行的SQL列表
	cache            *gcache.Cache                // 数据库缓存，包括底层连接池对象缓存及查询缓存；需要注意的是，事务查询不支持查询缓存
	schema           *gtype.String                // 手动切换的数据库名称
	tables           map[string]map[string]string // 数据库表结构
	logger           *glog.Logger                 // 日志管理对象
	maxIdleConnCount int                          // 连接池最大限制的连接数
	maxOpenConnCount int                          // 连接池最大打开的连接数
	maxConnLifetime  int                          // (单位秒)连接对象可重复使用的时间长度
}

// 执行的SQL对象
type Sql struct {
	Sql   string        // SQL语句(可能带有预处理占位符)
	Args  []interface{} // 预处理参数值列表
	Error error         // 执行结果(nil为成功)
	Start int64         // 执行开始时间(毫秒)
	End   int64         // 执行结束时间(毫秒)
}

// 表字段结构信息
type TableField struct {
	Index   int         // 用于字段排序(map类型是无序的)
	Name    string      // 字段名称
	Type    string      // 字段类型
	Null    bool        // 是否可为null
	Key     string      // 索引信息
	Default interface{} // 默认值
	Extra   string      // 其他信息
}

// 返回数据表记录值
type Value = *gvar.Var

// 返回数据表记录Map
type Record map[string]Value

// 返回数据表记录List
type Result []Record

// 关联数组，绑定一条数据表记录(使用别名)
type Map = map[string]interface{}

// 关联数组列表(索引从0开始的数组)，绑定多条记录(使用别名)
type List = []Map

const (
	OPTION_INSERT               = 0
	OPTION_REPLACE              = 1
	OPTION_SAVE                 = 2
	OPTION_IGNORE               = 3
	gDEFAULT_BATCH_NUM          = 10 // Per count for batch insert/replace/save
	gDEFAULT_CONN_MAX_LIFE_TIME = 30 // Max life time for per connection in pool.
)

var (
	// Instance map.
	instances = gmap.NewStrAnyMap(true)
)

// New creates ORM DB object with global configurations.
// The parameter <name> specifies the configuration group name,
// which is DEFAULT_GROUP_NAME in default.
func New(name ...string) (db DB, err error) {
	group := configs.defaultGroup
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}
	configs.RLock()
	defer configs.RUnlock()

	if len(configs.config) < 1 {
		return nil, errors.New("empty database configuration")
	}
	if _, ok := configs.config[group]; ok {
		if node, err := getConfigNodeByGroup(group, true); err == nil {
			base := &dbBase{
				group:           group,
				debug:           gtype.NewBool(),
				cache:           gcache.New(),
				schema:          gtype.NewString(),
				logger:          glog.Default(),
				maxConnLifetime: gDEFAULT_CONN_MAX_LIFE_TIME,
			}
			switch node.Type {
			case "mysql":
				base.db = &dbMysql{dbBase: base}
			case "pgsql":
				base.db = &dbPgsql{dbBase: base}
			case "mssql":
				base.db = &dbMssql{dbBase: base}
			case "sqlite":
				base.db = &dbSqlite{dbBase: base}
			case "oracle":
				base.db = &dbOracle{dbBase: base}
			default:
				return nil, errors.New(fmt.Sprintf(`unsupported database type "%s"`, node.Type))
			}
			return base.db, nil
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New(fmt.Sprintf("empty database configuration for item name '%s'", group))
	}
}

// Instance returns an instance for DB operations.
// The parameter <name> specifies the configuration group name,
// which is DEFAULT_GROUP_NAME in default.
func Instance(name ...string) (db DB, err error) {
	group := configs.defaultGroup
	if len(name) > 0 {
		group = name[0]
	}
	v := instances.GetOrSetFuncLock(group, func() interface{} {
		db, err = New(group)
		return db
	})
	if v != nil {
		return v.(DB), nil
	}
	return
}

// 获取指定数据库角色的一个配置项，内部根据权重计算负载均衡
func getConfigNodeByGroup(group string, master bool) (*ConfigNode, error) {
	if list, ok := configs.config[group]; ok {
		// 将master, slave集群列表拆分出来
		masterList := make(ConfigGroup, 0)
		slaveList := make(ConfigGroup, 0)
		for i := 0; i < len(list); i++ {
			if list[i].Role == "slave" {
				slaveList = append(slaveList, list[i])
			} else {
				masterList = append(masterList, list[i])
			}
		}
		if len(masterList) < 1 {
			return nil, errors.New("at least one master node configuration's need to make sense")
		}
		if len(slaveList) < 1 {
			slaveList = masterList
		}
		if master {
			return getConfigNodeByWeight(masterList), nil
		} else {
			return getConfigNodeByWeight(slaveList), nil
		}
	} else {
		return nil, errors.New(fmt.Sprintf("empty database configuration for item name '%s'", group))
	}
}

// 按照负载均衡算法(优先级配置)从数据库集群中选择一个配置节点出来使用
// 算法说明举例，
// 1、假如2个节点的priority都是1，那么随机大小范围为[0, 199]；
// 2、那么节点1的权重范围为[0, 99]，节点2的权重范围为[100, 199]，比例为1:1；
// 3、假如计算出的随机数为99;
// 4、那么选择的配置为节点1;
func getConfigNodeByWeight(cg ConfigGroup) *ConfigNode {
	if len(cg) < 2 {
		return &cg[0]
	}
	var total int
	for i := 0; i < len(cg); i++ {
		total += cg[i].Weight * 100
	}
	// 如果total为0表示所有连接都没有配置priority属性，那么默认都是1
	if total == 0 {
		for i := 0; i < len(cg); i++ {
			cg[i].Weight = 1
			total += cg[i].Weight * 100
		}
	}
	// 不能取到末尾的边界点
	r := grand.N(0, total)
	if r > 0 {
		r -= 1
	}
	min := 0
	max := 0
	for i := 0; i < len(cg); i++ {
		max = min + cg[i].Weight*100
		//fmt.Printf("r: %d, min: %d, max: %d\n", r, min, max)
		if r >= min && r < max {
			return &cg[i]
		} else {
			min = max
		}
	}
	return nil
}

// 获得底层数据库链接对象
func (bs *dbBase) getSqlDb(master bool) (sqlDb *sql.DB, err error) {
	// 负载均衡
	node, err := getConfigNodeByGroup(bs.group, master)
	if err != nil {
		return nil, err
	}
	// 默认值设定
	if node.Charset == "" {
		node.Charset = "utf8"
	}
	// 缓存连接对象(该对象其实是一个连接池对象)
	v := bs.cache.GetOrSetFuncLock(node.String(), func() interface{} {
		sqlDb, err = bs.db.Open(node)
		if err != nil {
			return nil
		}
		// 接口对象可能会覆盖这些连接参数，所以这里优先判断有误设置连接池属性。
		// 若无设置则使用配置节点的连接池参数
		if bs.maxIdleConnCount > 0 {
			sqlDb.SetMaxIdleConns(bs.maxIdleConnCount)
		} else if node.MaxIdleConnCount > 0 {
			sqlDb.SetMaxIdleConns(node.MaxIdleConnCount)
		}

		if bs.maxOpenConnCount > 0 {
			sqlDb.SetMaxOpenConns(bs.maxOpenConnCount)
		} else if node.MaxOpenConnCount > 0 {
			sqlDb.SetMaxOpenConns(node.MaxOpenConnCount)
		}

		if bs.maxConnLifetime > 0 {
			sqlDb.SetConnMaxLifetime(time.Duration(bs.maxConnLifetime) * time.Second)
		} else if node.MaxConnLifetime > 0 {
			sqlDb.SetConnMaxLifetime(time.Duration(node.MaxConnLifetime) * time.Second)
		}
		return sqlDb
	}, 0)
	if v != nil && sqlDb == nil {
		sqlDb = v.(*sql.DB)
	}
	// 是否开启调试模式
	if node.Debug {
		bs.db.SetDebug(node.Debug)
	}
	// 是否手动选择数据库
	if v := bs.schema.Val(); v != "" {
		if e := bs.db.setSchema(sqlDb, v); e != nil {
			err = e
		}
	}
	return
}

// 切换当前数据库对象操作的数据库。
func (bs *dbBase) SetSchema(schema string) {
	bs.schema.Set(schema)
}

// 创建底层数据库master链接对象。
func (bs *dbBase) Master() (*sql.DB, error) {
	return bs.getSqlDb(true)
}

// 创建底层数据库slave链接对象。
func (bs *dbBase) Slave() (*sql.DB, error) {
	return bs.getSqlDb(false)
}
