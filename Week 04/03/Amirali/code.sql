-- Section1
SELECT name, phone FROM customers INNER JOIN orders ON customers.id = orders.customer_id GROUP BY customer_id ORDER BY COUNT(*) DESC LIMIT 0,1;
-- Section2
SELECT f.id AS id, f.name AS name FROM foods f INNER JOIN restaurant_foods r ON f.id = r.food_id INNER JOIN orders o ON r.id = o.restaurant_food_id GROUP BY 1 ORDER BY AVG(o.rate) DESC, f.id LIMIT 10;
-- Section3
SELECT r.id, r.name FROM restaurants r INNER JOIN restaurant_foods f ON r.id = f.restaurant_id INNER JOIN orders o ON f.id = o.restaurant_food_id GROUP BY 1 ORDER BY AVG(o.rate) DESC, 1 LIMIT 10;
-- Section4
SELECT c.name AS name, c.phone AS phone FROM customers c INNER JOIN orders o ON c.id = o.customer_id INNER JOIN restaurant_foods f ON o.restaurant_food_id = f.id GROUP BY c.id HAVING COUNT(DISTINCT(f.restaurant_id)) >=5 ORDER BY 1;