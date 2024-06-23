CREATE TABLE IF NOT EXISTS users(
                                    id SERIAL PRIMARY KEY,
                                    name TEXT NOT NULL,
                                    email TEXT NOT NULL,
                                    phone TEXT NOT NULL,
                                    password TEXT NOT NULL,
                                    dob DATE NOT NULL,
                                    role TEXT NOT NULL,
                                    created_at TIMESTAMP NOT NULL,
                                    updated_at TIMESTAMP NOT NULL
);
