
local M = {}
M.array = {}


function M.bsort()
    local array = M.array
    local len = #array
    for i = 1,len do
        for j = 1, len - 1 do
           if array[j] > array[j + 1] then
                local temp = array[j]
                array[j] = array[j + 1]
                array[j + 1] = temp
           end
        end
    end
    return array
end


return M
