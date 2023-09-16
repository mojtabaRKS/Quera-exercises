-- Section1
CREATE INDEX idx1 ON orders (created_at, total);
-- Section2
CREATE INDEX idx2 ON orders (user_id, created_at, total);
--section3

