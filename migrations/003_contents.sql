CREATE TABLE categories (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            name TEXT NOT NULL UNIQUE
);

CREATE TABLE tags (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      name TEXT NOT NULL UNIQUE
);

CREATE TABLE contents (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          title TEXT NOT NULL,
                          body TEXT NOT NULL,
                          image TEXT,
                          category_id UUID REFERENCES categories(id),
                          author_id UUID REFERENCES users(id),
                          created_at TIMESTAMP DEFAULT now(),
                          updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE content_tags (
                              content_id UUID REFERENCES contents(id) ON DELETE CASCADE,
                              tag_id UUID REFERENCES tags(id) ON DELETE CASCADE,
                              PRIMARY KEY (content_id, tag_id)
);
