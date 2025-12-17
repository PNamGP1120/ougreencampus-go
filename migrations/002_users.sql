CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       name TEXT NOT NULL,
                       avatar TEXT,
                       role user_role NOT NULL DEFAULT 'student',
                       status user_status NOT NULL DEFAULT 'active',
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now()
);
