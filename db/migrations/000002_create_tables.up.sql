CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY, username VARCHAR(255) NOT NULL, created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

CREATE TABLE categories (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY, title VARCHAR(255) NOT NULL, description TEXT, created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

CREATE TABLE todos (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY, user_id UUID NOT NULL, heading VARCHAR(255) NOT NULL, body TEXT, category_id UUID, created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, FOREIGN KEY (user_id) REFERENCES users (id), FOREIGN KEY (category_id) REFERENCES categories (id)
);

CREATE TABLE reminders (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY, todo_id UUID NOT NULL, snooze BOOLEAN NOT NULL, snooze_timeout INTERVAL, remind_time TIMESTAMPTZ NOT NULL, description TEXT, created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp, FOREIGN KEY (todo_id) REFERENCES todos (id)
);