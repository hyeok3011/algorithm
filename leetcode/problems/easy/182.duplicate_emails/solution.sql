-- Runtine Runtime 858 ms
SELECT email AS Email FROM Person GROUP BY email HAVING COUNT(*) > 1;

-- Runtime 740 ms
SELECT p1.email AS Email
FROM Person p1
INNER JOIN Person p2 ON p1.email = p2.email AND p1.id <> p2.id
group by Email

-- 다른사람의 솔류션 단순 중복제거를 위함이면 group by보다 DISTINCT 더 효율적일듯 하다
-- Runtime 700 ms
SELECT DISTINCT(p1.email) from 
Person p1 JOIN Person p2 ON
p1.email = p2.email AND p1.id <> p2.id;