

-- leetcode 196
-- delete rows with the same email, keep the one only with smallest id
DELETE p1
FROM Person p1 LEFT JOIN 
(
	SELECT min(id) as min_id
	FROM Person 
	GROUP BY email
) p2 ON p1.id = p2.min_id
WHERE p2.min_id IS NULL

-- or 
DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id

