CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES users(id),
    name TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS feedbacks (
    id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES companies(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);