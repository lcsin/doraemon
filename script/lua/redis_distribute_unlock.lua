local key = KEYS[1]
local uniqueID = ARGV[1]

if redis.call("get", key) == uniqueID then
    redis.call("del", key)
    return 1
end
return -1
