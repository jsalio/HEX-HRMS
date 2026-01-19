-- PostgreSQL Initialization Script
-- This script runs automatically on first container startup
-- Reference: https://sliplane.io/blog/best-practices-for-postgres-in-docker

-- Enable pg_stat_statements extension for query performance monitoring
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

-- Enable uuid-ossp for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Log successful initialization
DO $$
BEGIN
    RAISE NOTICE 'HRMS database initialized successfully';
END $$;