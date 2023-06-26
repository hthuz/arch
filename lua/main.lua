
local array = {5,11,3,4,2,1,5}

local res = require("sort").bsort(array)
for k,v in pairs(res) do
    print(k,v)
end


