-- Section1
CREATE INDEX orders_created_at_totalx ON orders (created_at, total);
-- Section2
CREATE INDEX orders_user_id_created_at_totalx ON orders (user_id, created_at, total);
-- Section3
WITH RECURSIVE DateRange AS (
    SELECT DATE('2020-01-01') AS date
    UNION ALL
    SELECT DATE_ADD(date, INTERVAL 1 DAY)
    FROM DateRange
    WHERE DATE_ADD(date, INTERVAL 1 DAY) <= '2021-12-11'
)
SELECT
    DateRange.date AS date,
    IFNULL(SUM(orders.total), 0) AS total
FROM DateRange
LEFT JOIN orders ON DateRange.date = DATE(orders.created_at)
GROUP BY DateRange.date
ORDER BY DateRange.date;


