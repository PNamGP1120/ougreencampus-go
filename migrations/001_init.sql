CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_role AS ENUM ('guest', 'student', 'organizer', 'admin');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned');

CREATE TYPE activity_type AS ENUM ('program', 'contest', 'campaign');
CREATE TYPE activity_status AS ENUM ('draft', 'active', 'closed');

CREATE TYPE event_status AS ENUM ('upcoming', 'ongoing', 'finished');
