# Write your MySQL query statement below
# https://leetcode.com/problems/employees-earning-more-than-their-managers/description/
# Runtime 758 ms Beats 77.63%
SELECT e.name AS Employee FROM employee e
INNER JOIN employee m on e.managerId = m.id
where e.salary > m.salary