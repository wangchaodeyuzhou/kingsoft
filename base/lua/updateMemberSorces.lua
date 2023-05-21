local FIND, ID, RANK, SCORE, PRERANK = 1, 2, 3, 4, 5
local leaderboard = KEYS[1]
local res = {}
for i, v in pairs(ARGV) do
    local data = cjson.decode(v)
    local info = {}
    info[FIND] = 0
    info[ID] = data["Member"]
    info[SCORE] = data["Score"]
    local rank = redis.call("ZRANK", leaderboard, data["Member"])
    if rank == false then
        info[PRERANK] = 0
    else
        info[PRERANK] = rank
    end
    redis.call("ZADD", leaderboard, data["Score"], data["Member"])
    res[i] = info
end
for i, v in pairs(res) do
    local rank = redis.call("ZRANK", leaderboard, v[ID])
    v[RANK] = rank
end
return res