-- 000001_roles.up.sql
-- Create enum type for user types
CREATE TYPE user_type AS ENUM (
    'ADMIN',
    'USER',
    'MODERATOR',
    'GUEST'
);

-- Create roles table
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_type user_type NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index on user_type
CREATE INDEX idx_roles_user_type ON roles(user_type);

-- Create trigger for updating timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_roles_updated_at
    BEFORE UPDATE ON roles
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Insert default roles
INSERT INTO roles (user_type) VALUES
    ('ADMIN'),
    ('USER'),
    ('MODERATOR'),
    ('GUEST')
ON CONFLICT DO NOTHING;
