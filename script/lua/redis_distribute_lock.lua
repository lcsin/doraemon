local key = KEYS[1]
local uniqueID = ARGV[1]
local ttl = tonumber(ARGV[2])

if redis.call("setnx", key, uniqueID) == 1 then
    redis.call("expire", key, ttl)
    return 1
elseif redis.call("get", key) == uniqueID then
    redis.call("expire", key, ttl)
    return 1
end
return -1

