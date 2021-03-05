local mysql  = require("resty/mysql")

InitDatabaseContext = { _version = "1.0.1",host,port,user,password,database,db }

-- 连接初始化
function InitDatabaseContext:new(o,host,port,user,password,database)
  local o = o or {}
  setmetatable(o,self)
  self.__index = self
  self.host = host or "localhost"
  self.port = port or 3306
  self.user = user or "user"
  self.password = password or "password"
  self.database = database or "database"

  -- 创建 mysql 实例 
  local db,err = mysql:new()
  if not db then
    return nil,string.format("new mysql error: %s",err)
  end

  -- 设置超时时间
  db:set_timeout(1000)

  -- 创建连接
  local res,err,errno,sqlstate =  db:connect({
    host = self.host,
    port = self.port,
    user = self.user,
    password = self.password,
    database = self.database,
  })
  if not res then
    self:close_database()
    return nil,string.format("connect mysql error: %s,errno: %s,sqlstate:%s",err,errno,sqlstate)
  end
 
  self.db = db
  return o,""
end

-- 连接关闭
function InitDatabaseContext:close_database()
  if not self.db then
    return
  end
  self.db:close()
end

-- 数据库初始化
function InitDatabaseContext:check(ngx)
  assert(self.db ~= nil,"db is not initialized")
  -- 初始化数据库
  local create_db_sql = string.format("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci",self.database)
  local res,err,errno,sqlstate = self.db:query(create_db_sql)
  if not res then
    self:close_database()
    return false,string.format("init database error: %s,errno: %s,sqlstate:%s",err,errno,sqlstate)
  end
 
  -- 初始化表
  local res,err,errno,sqlstate = self.db:query([[
  CREATE TABLE IF NOT EXISTS `account` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账号ID',
      `account` varchar(50) NOT NULL COMMENT '账户',
      `password` varchar(50) NOT NULL COMMENT '密码',
      `create_time` bigint(20) NOT NULL COMMENT '创建时间',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号';
  ]])
  if not res then
    self:close_database()
    return false, string.format("init table[account] error: %s,errno: %s,sqlstate:%s",err,errno,sqlstate)
  end
  
  res,err,errno,sqlstate = self.db:query([[
  CREATE TABLE IF NOT EXISTS `invite_code` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '邀请码ID',
      `invite_code` varchar(50) NOT NULL COMMENT '邀请码',
      `used` smallint NOT NULL COMMENT '是否已使用',
      `create_time` bigint(20) NOT NULL COMMENT '创建时间',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邀请码';
  ]])
  if not res then
    self.db:close()
    return false,string.format("init table[invite_code] error: %s,errno: %s,sqlstate:%s",err,errno,sqlstate)
  end
  
  self:close_database()
  return true,""
end


-- openresty 调用
local r,err = InitDatabaseContext:new(nil,"127.0.0.1",3306,"user","password","auth")
if not r then
  ngx.log(ngx.ERR,err)
  ngx.exit(ngx.HTTP_INTERNAL_SERVER_ERROR)
end

local res,err = r:check(ngx)
if not res then
  ngx.log(ngx.ERR,err)
  ngx.exit(ngx.HTTP_INTERNAL_SERVER_ERROR)
end

ngx.exit(ngx.HTTP_OK)


