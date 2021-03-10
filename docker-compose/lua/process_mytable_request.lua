local mytable  = require("mytable")

local function filter_method(ngx,allow_method)
 if ngx.req.get_method() ~= allow_method then
   ngx.exit(ngx.HTTP_NOT_ALLOWED)
 end
end

local function check(ngx,val,message)
  if not val then
    ngx.say(cjson.encode({message = message}))
    ngx.exit(ngx.HTTP_OK)
  end
end

-- 处理 nginx 请求
local db_host = os.getenv("DB_HOST")
local db_port = os.getenv("DB_PORT")
local db_user = os.getenv("DB_USER")
local db_password = os.getenv("DB_PASSWORD")
local db_lib = os.getenv("DB_LIB")
local g_mytable

if g_mytable == nil then
  local ok,mytable_temp = mytable:new(nil,db_host,db_port,db_user,db_password,db_lib)
  if not ok then
    error("mytable new error")
  end
  g_mytable = mytable_temp
end

if ngx.var.uri == "/mytable/query_table_list" then
  filter_method(ngx,"GET")

  ngx.say(cjson.encode(g_mytable:query_table_list()))

elseif ngx.var.uri == "/mytable/query_table_struct" then
  filter_method(ngx,"GET")
  
  local table_name = ngx.var.arg_table_name
  check(ngx,table_name ~= nil,"param[table_name] is empty")
  
  ngx.say(cjson.encode(g_mytable:query_table_struct(table_name)))
  
elseif ngx.var.uri == "/mytable/query_table_data" then
  filter_method(ngx,"GET")

  local table_name = ngx.var.arg_table_name
  check(ngx,table_name ~= nil,"param[table_name] is empty")
  
  ngx.say(cjson.encode(g_mytable:query_table_data(table_name)))

elseif ngx.var.uri == "/mytable/update_table_data" then
  filter_method(ngx,"POST") 
  
  ngx.req.read_body()
  local body = ngx.req.get_body_data()
  check(ngx,body ~= nil,"request body is empty")
  
  local param = cjson.decode(ngx.req.get_body_data())
  local table_name = param["table_name"]
  local rows = param["rows"]
  check(ngx,table_name ~= nil,"param[table_name] is empty")
  check(ngx,rows ~= nil,"param[rows] is empty")
  
  g_mytable:update_table_data(table_name,rows)

elseif ngx.var.uri == "/mytable/delete_table_data" then
  filter_method(ngx,"POST") 

  ngx.req.read_body()
  local body = ngx.req.get_body_data()
  check(ngx,body ~= nil,"request body is empty")

  local param = cjson.decode(ngx.req.get_body_data())
  local table_name = param["table_name"]
  local rows = param["rows"]
  check(ngx,table_name ~= nil,"param[table_name] is empty")
  check(ngx,rows ~= nil,"param[rows] is empty")
  
  g_mytable:delete_table_data(table_name,rows)

elseif ngx.var.uri == "/mytable/insert_table_data" then
  filter_method(ngx,"POST") 

  ngx.req.read_body()
  local body = ngx.req.get_body_data()
  check(ngx,body ~= nil,"request body is empty")
  
  local param = cjson.decode(ngx.req.get_body_data())
  local table_name = param["table_name"]
  local rows = param["rows"]
  check(ngx,table_name ~= nil,"param[table_name] is empty")
  check(ngx,rows ~= nil,"param[rows] is empty")
  
  g_mytable:insert_table_data(table_name,rows)

else
  ngx.exit(ngx.HTTP_METHOD_NOT_IMPLEMENTED)
end

return _M
