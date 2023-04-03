

local M = {}

-- This function is invisible outside this file
local function sayMyname()
    print("Autentico")
end


-- You can run this function
function M.sayHello()
    print('Hello, ')
    sayMyname()
end

return M

