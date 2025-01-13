-- Connect to PostgreSQL in Docker
docker exec -it commerce-hub-go-pg psql -U admin -d commerce_hub

-- Show all databases
\l

-- Create a new database
-- CREATE DATABASE commerce_hub;

-- Connect to the new database
\c commerce_hub

-- Create an admin user
-- CREATE USER admin WITH PASSWORD 'securepassword123';
-- ALTER USER admin WITH SUPERUSER;

-- Create the 'customers' table
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(128),
    last_name VARCHAR(128),
    company_name VARCHAR(256),
    customer_type INT NOT NULL,
    -- contact_info JSONB
);

\dt

\d+ customers

\q