-- Initialize PostgreSQL database for Market Account Microservices
-- This script sets up the basic database structure and schemas

-- Create databases for each service if they don't exist
CREATE DATABASE IF NOT EXISTS marketplace;

-- Switch to marketplace database
\c marketplace;

-- Create schemas for different services
CREATE SCHEMA IF NOT EXISTS user_service;
CREATE SCHEMA IF NOT EXISTS product_service;
CREATE SCHEMA IF NOT EXISTS payment_service;

-- Create extensions that might be needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- User Service Tables
CREATE TABLE IF NOT EXISTS user_service.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) DEFAULT 'customer',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Product Service Tables (placeholder)
CREATE TABLE IF NOT EXISTS product_service.products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INTEGER DEFAULT 0,
    category VARCHAR(100),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Payment Service Tables (placeholder)
CREATE TABLE IF NOT EXISTS payment_service.payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) DEFAULT 'VND',
    status VARCHAR(50) DEFAULT 'pending',
    payment_method VARCHAR(100),
    transaction_id VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_users_email ON user_service.users(email);
CREATE INDEX IF NOT EXISTS idx_users_role ON user_service.users(role);
CREATE INDEX IF NOT EXISTS idx_products_category ON product_service.products(category);
CREATE INDEX IF NOT EXISTS idx_products_active ON product_service.products(is_active);
CREATE INDEX IF NOT EXISTS idx_payments_user ON payment_service.payments(user_id);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payment_service.payments(status);

-- Insert sample data for development
INSERT INTO user_service.users (email, password_hash, first_name, last_name, role) 
VALUES 
    ('admin@marketplace.com', crypt('admin123', gen_salt('bf')), 'Admin', 'User', 'admin'),
    ('user@marketplace.com', crypt('user123', gen_salt('bf')), 'Regular', 'User', 'customer')
ON CONFLICT (email) DO NOTHING;

COMMIT; 