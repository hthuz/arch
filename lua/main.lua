
local function tprint(table)
    for k,v in pairs(table) do
        print(k,v)
    end
end

-- Use require to run the file mod.lua
-- Note that if you require this mod multiple times, it is only run once
local mod = require('mod')

mod.sayHello()
print(mod)

for k,v in pairs(mod) do
    print(k,v)
end


local t = {}
t[1] = 12
t[2] = 34
t[4] = 33

tprint(t)

