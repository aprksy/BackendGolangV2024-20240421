-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

CREATE TABLE estate (
  id uuid PRIMARY KEY,
  length INTEGER,
  width INTEGER
);

CREATE TABLE plot (
  id SERIAL PRIMARY KEY,
  estate_id uuid REFERENCES estate(id) ON DELETE CASCADE,
  x INTEGER,
  y INTEGER,
  height INTEGER
);

CREATE VIEW estate_stats AS
SELECT 
  e.id AS estate_id,
  COUNT(*) AS count,
  MIN(p.height) AS min,
  MAX(p.height) AS max,
  PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY p.height) AS median
FROM estate e
LEFT JOIN plot p ON p.estate_id = e.id
GROUP BY e.id;
