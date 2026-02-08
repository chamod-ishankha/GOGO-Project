-- Create schema if not exists 
CREATE SCHEMA IF NOT EXISTS gogo;

-- Create users table
CREATE TABLE IF NOT EXISTS gogo.users (
    id BIGSERIAL PRIMARY KEY,          -- Auto-incrementing primary key, better as BIGSERIAL for production
    name VARCHAR(100) NOT NULL,        -- User full name
    email VARCHAR(150) UNIQUE NOT NULL, -- Unique email for login
    password VARCHAR(255) NOT NULL,    -- Hashed password
    role VARCHAR(20) NOT NULL DEFAULT 'rider', -- user roles: rider / driver / admin
    created_at TIMESTAMPTZ DEFAULT NOW(), -- timestamp with timezone
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS gogo.drivers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE,
    license_number VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    is_available BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_driver_user
        FOREIGN KEY (user_id)
        REFERENCES gogo.users(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS gogo.vehicles (
    id BIGSERIAL PRIMARY KEY,
    driver_id BIGINT NOT NULL,
    vehicle_type VARCHAR(20) NOT NULL, -- car, bike, tuk
    make VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    color VARCHAR(30),
    plate_number VARCHAR(20) UNIQUE NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_vehicle_driver
        FOREIGN KEY (driver_id)
        REFERENCES gogo.drivers(id)
        ON DELETE CASCADE
);


