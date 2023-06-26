
require("sort").array = {412,412,12,34,123,6634,14}

local res = require("sort").bsort()
for k,v in pairs(res) do
    print(k,v)
end


