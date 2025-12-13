CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role TEXT NOT NULL,
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE contents (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          title TEXT NOT NULL,
                          body TEXT,
                          type TEXT NOT NULL,
                          status TEXT DEFAULT 'draft',
                          author_id UUID,
                          created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE events (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        title TEXT NOT NULL,
                        description TEXT,
                        start_time TIMESTAMP,
                        end_time TIMESTAMP,
                        location TEXT,
                        capacity INT,
                        created_by UUID,
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE event_registrations (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     event_id UUID,
                                     user_id UUID,
                                     checked_in BOOLEAN DEFAULT FALSE,
                                     created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE activities (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name TEXT NOT NULL,
                            description TEXT,
                            type TEXT,
                            start_date TIMESTAMP,
                            end_date TIMESTAMP,
                            created_by UUID,
                            created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE submissions (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             activity_id UUID,
                             user_id UUID,
                             content TEXT,
                             score INT,
                             status TEXT,
                             created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE system_configs (
                                key TEXT PRIMARY KEY,
                                value TEXT
);

CREATE TABLE audit_logs (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            actor_id UUID,
                            action TEXT,
                            target TEXT,
                            created_at TIMESTAMP DEFAULT NOW()
);
