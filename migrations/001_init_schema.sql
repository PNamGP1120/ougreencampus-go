CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- USERS
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role TEXT NOT NULL,
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now(),
                       deleted_at TIMESTAMP
);

-- CONTENTS
CREATE TABLE contents (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          title TEXT NOT NULL,
                          body TEXT,
                          type TEXT NOT NULL,
                          is_featured BOOLEAN DEFAULT FALSE,
                          author_id UUID,
                          created_at TIMESTAMP DEFAULT now(),
                          updated_at TIMESTAMP DEFAULT now(),
                          deleted_at TIMESTAMP
);

-- EVENTS
CREATE TABLE events (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        title TEXT,
                        description TEXT,
                        start_time TIMESTAMP,
                        end_time TIMESTAMP,
                        location TEXT,
                        capacity INT,
                        created_by UUID,
                        created_at TIMESTAMP DEFAULT now(),
                        updated_at TIMESTAMP DEFAULT now(),
                        deleted_at TIMESTAMP
);

CREATE TABLE event_registrations (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     event_id UUID,
                                     user_id UUID,
                                     checked_in BOOLEAN DEFAULT FALSE,
                                     created_at TIMESTAMP DEFAULT now()
);

-- ACTIVITIES
CREATE TABLE activities (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name TEXT NOT NULL,
                            description TEXT,
                            type TEXT NOT NULL,
                            status TEXT DEFAULT 'draft',
                            start_at TIMESTAMP,
                            end_at TIMESTAMP,
                            created_by UUID,
                            created_at TIMESTAMP DEFAULT now(),
                            updated_at TIMESTAMP DEFAULT now(),
                            deleted_at TIMESTAMP
);

CREATE TABLE submissions (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             activity_id UUID,
                             user_id UUID,
                             content TEXT,
                             status TEXT,
                             score INT,
                             feedback TEXT,
                             reviewed_by UUID,
                             reviewed_at TIMESTAMP,
                             created_at TIMESTAMP DEFAULT now(),
                             updated_at TIMESTAMP DEFAULT now(),
                             deleted_at TIMESTAMP
);

CREATE TABLE campaign_tasks (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                activity_id UUID,
                                title TEXT,
                                detail TEXT,
                                points INT DEFAULT 0,
                                due_at TIMESTAMP,
                                is_active BOOLEAN DEFAULT TRUE,
                                created_at TIMESTAMP DEFAULT now(),
                                updated_at TIMESTAMP DEFAULT now(),
                                deleted_at TIMESTAMP
);

CREATE TABLE campaign_progresses (
                                     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     activity_id UUID,
                                     task_id UUID,
                                     user_id UUID,
                                     status TEXT,
                                     note TEXT,
                                     evidence JSONB,
                                     reviewed_by UUID,
                                     reviewed_at TIMESTAMP,
                                     created_at TIMESTAMP DEFAULT now(),
                                     updated_at TIMESTAMP DEFAULT now(),
                                     deleted_at TIMESTAMP
);

CREATE TABLE impact_metrics (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                activity_id UUID,
                                key TEXT,
                                value DOUBLE PRECISION,
                                unit TEXT,
                                note TEXT,
                                updated_by UUID,
                                updated_at TIMESTAMP DEFAULT now()
);

-- SYSTEM
CREATE TABLE system_configs (
                                key TEXT PRIMARY KEY,
                                value TEXT,
                                updated_by UUID,
                                updated_at TIMESTAMP DEFAULT now(),
                                deleted_at TIMESTAMP
);

CREATE TABLE audit_logs (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            actor_id UUID,
                            role TEXT,
                            action TEXT,
                            entity TEXT,
                            entity_id TEXT,
                            metadata JSONB,
                            ip TEXT,
                            user_agent TEXT,
                            created_at TIMESTAMP DEFAULT now()
);
