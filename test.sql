-- Section1
SELECT c.name, c.phone FROM customers c JOIN orders o ON c.id = o.customer_id GROUP BY c.name, c.phone ORDER BY count(o.id)  DESC LIMIT 1;
-- Section2
SELECT f.id, f.name
FROM foods f
JOIN (
       SELECT rf.food_id, AVG(o.rate) AS avg_rate
       FROM restaurant_foods rf
       JOIN orders o ON rf.id = o.restaurant_food_id
         GROUP BY rf.food_id
         ORDER BY avg_rate DESC, rf.food_id ASC
	 LIMIT 10)
	 AS top_foods ON f.id = top_foods.food_id;
-- Section3
SELECT r.id, r.name 
FROM restaurants r 
JOIN (
	SELECT rf.restaurant_id, AVG(o.rate) AS avg_rate 
	FROM restaurant_foods rf     
	JOIN orders o ON rf.id = o.restaurant_food_id     
	GROUP BY rf.restaurant_id     
	ORDER BY avg_rate DESC, rf.restaurant_id ASC     
	LIMIT 10) 
	AS top_restaurants ON r.id = top_restaurants.restaurant_id;
-- Section4
SELECT c.name, c.phone 
FROM customers c 
JOIN (
	SELECT customer_id     
	FROM (SELECT o.customer_id, COUNT(DISTINCT rf.restaurant_id) 
	AS num_restaurants         
	FROM orders o
        JOIN restaurant_foods rf ON o.restaurant_food_id = rf.id   
        GROUP BY o.customer_id         
        HAVING COUNT(DISTINCT rf.restaurant_id) >= 5 ) AS frequent_customers ) AS top_customers
        ON c.id = top_customers.customer_id 
        ORDER BY c.name ASC;


