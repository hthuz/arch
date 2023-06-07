

-- Use require to run the file mod.lua
-- Note that if you require this mod multiple times, it is only run once
local mod = require('mod')

mod.sayHello()
print(mod)

for k,v in pairs(mod) do
    print(k,v)
end





