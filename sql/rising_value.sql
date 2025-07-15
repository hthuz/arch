

-- leetcode 197
-- find rows with increasing values (rows with value higher than prev low)

SELECT w2.id
FROM Weather w1 JOIN Weather w2 ON w1.recordDate + 1 = w2.recordDate
WHERE w2.Temperature -	w1.Temperature > 0
