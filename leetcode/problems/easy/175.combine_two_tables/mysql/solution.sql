# Write your MySQL query statement below
-- https://leetcode.com/problems/combine-two-tables/description/
-- Runtime 644 ms Beats 99.49% Memory 100%
SELECT p.firstName, p.lastName, a.city, a.state FROM person AS p 
LEFT JOIN Address a ON a.personId = p.personId;