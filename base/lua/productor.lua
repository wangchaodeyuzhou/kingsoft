---
--- Generated by Luanalysis
--- Created by Administrator.
--- DateTime: 2022/11/16 16:38
---

local newProductor

function productor()
    local i = 0
    while true do
        i = i + 1
        send(i)
    end
end

function consumer()
    while true do
        local i = recevice()
        print(i)
    end
end

function recevice()
    local status, value = coroutine.resume(newProductor)
    return value
end

function send(x)
    coroutine.yield(x)
end

newProductor = coroutine.create(productor)
consumer()