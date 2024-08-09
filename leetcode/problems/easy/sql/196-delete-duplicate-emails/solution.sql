-- https://leetcode.com/problems/delete-duplicate-emails/description/
DELETE p1 from Person p1
INNER JOIN Person p2 ON p1.email = p2.email 
WHERE p1.id >  p2.id