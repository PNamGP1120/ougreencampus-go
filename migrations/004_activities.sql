CREATE TABLE activities (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            title TEXT NOT NULL,
                            description TEXT,
                            image TEXT,
                            type activity_type NOT NULL,
                            status activity_status NOT NULL DEFAULT 'draft',
                            created_by UUID REFERENCES users(id),
                            created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE activity_participants (
                                       activity_id UUID REFERENCES activities(id) ON DELETE CASCADE,
                                       user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                                       joined_at TIMESTAMP DEFAULT now(),
                                       PRIMARY KEY (activity_id, user_id)
);

CREATE TABLE submissions (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             activity_id UUID REFERENCES activities(id),
                             user_id UUID REFERENCES users(id),
                             content TEXT,
                             file_url TEXT,
                             score INT,
                             comment TEXT,
                             created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE campaign_tasks (
                                id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                activity_id UUID REFERENCES activities(id),
                                title TEXT NOT NULL,
                                target INT NOT NULL
);

CREATE TABLE task_progress (
                               task_id UUID REFERENCES campaign_tasks(id),
                               user_id UUID REFERENCES users(id),
                               value INT NOT NULL,
                               updated_at TIMESTAMP DEFAULT now(),
                               PRIMARY KEY (task_id, user_id)
);
