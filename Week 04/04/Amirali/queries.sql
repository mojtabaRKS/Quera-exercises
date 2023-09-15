-- Section1
ALTER TABLE orders ADD INDEX indx1(created_at, total);
-- Section2
ALTER TABLE orders ADD INDEX indx2(user_id, created_at, total);
-- Section3
WITH RECURSIVE t AS (SELECT MIN(DATE(created_at)) AS dt FROM orders UNION SELECT DATE_ADD(t.dt, INTERVAL 1 DAY) FROM t WHERE DATE_ADD(t.dt, INTERVAL 1 DAY) <= (SELECT MAX(DATE(created_at)) FROM orders)) SELECT dt, SUM(COALESCE(orders.total, 0)) FROM t LEFT JOIN orders ON t.dt = DATE(orders.created_at) GROUP BY 1 ORDER BY 1;