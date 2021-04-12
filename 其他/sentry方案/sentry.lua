local redis = require "resty.redis"
local cjson = require "cjson"

local r = redis:new()
local ok, err = r:connect("10.10.134.128", "6379")
r:select(6)
if not ok then
  return
end

local sha1 = "f11055f624b52dc2c757668967e525df25dd7f15"

--local uri_args = ngx.req.get_uri_args()
local body = ngx.req.get_body_data()
local header = ngx.req.get_headers()
local request_uri = ngx.var.request_uri

now = ngx.var.msec..math.random(1000000000, os.time())
--时间戳作为obj的一个参数，并作为body的key
local obj = {
        uri = request_uri,
        timestamp = now,
	    header = header
}

local sentryInfo = header["x-sentry-auth"]

--构造sentry key
sentry_key = "sentry:header:uri"

key = string.match(sentryInfo,"sentry_key=.+.,")
secret = string.match(sentryInfo,"sentry_secret=.+.")

if key ~= nil and secret ~= nil then
	k = string.sub(key,string.find(key,"=",1) + 1,string.len(key)-1)
	s = string.sub(secret,string.find(secret,"=",1) + 1,string.len(secret))
	sentry_key = "sentry:"..k..":"..s
end

--增加限流，每60秒最多10000个
limit_key = sentry_key..":".."limit"
limit = r:evalsha(sha1,1,limit_key,10000,60)

if limit == 0 then
    ngx.say('访问过于频繁')
    return
end

--header、uri存入redis
jsonData = cjson.encode(obj)
ok,err = r:lpush(sentry_key,jsonData)
if not ok then
        ngx.say("set data error",err,obj)
        return
end
--sentry_key存入列表
ok,err = r:sadd("sentry_keys",sentry_key)
if not ok then
    ngx.say("add sentry_key error",err,obj)
	return
end
--body存入redis
body_key = "sentry:body:"..now
ok,err = r:set(body_key,body)
if not ok then
	ngx.say("set body error: ",err,body_key)
end
ngx.say("success")