-- +goose Up
-- Apply database schema changes

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    status VARCHAR(30) CHECK (status IN ('pending', 'approved', 'rejected')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    address VARCHAR(100) NOT NULL,
    status SMALLINT NOT NULL, -- 0: inactive, 1: active
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    department_id INT,
    email VARCHAR(45) NOT NULL,
    password TEXT NOT NULL,
    name VARCHAR(45) NOT NULL,
    surname VARCHAR(45) NOT NULL,
    gender VARCHAR(20) NOT NULL,
    dob DATE NOT NULL,
    mobile VARCHAR(15) NOT NULL,
    country_id INT NOT NULL,
    resident_country_id INT NOT NULL,
    avatar VARCHAR(100),
    verification_status SMALLINT DEFAULT 0, -- 0: unverified, 1: verified
    status SMALLINT NOT NULL, -- 0: inactive, 1: active
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT fk_users_countries FOREIGN KEY (country_id) REFERENCES countries (id),
    CONSTRAINT fk_users_resident_countries FOREIGN KEY (resident_country_id) REFERENCES countries (id),
    CONSTRAINT fk_users_departments FOREIGN KEY (department_id) REFERENCES departments (id)
);

CREATE TABLE IF NOT EXISTS volunteer_details (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    department_id INT NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_volunteer_details_users FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_volunteer_details_departments FOREIGN KEY (department_id) REFERENCES departments (id)
);

CREATE TABLE IF NOT EXISTS requests (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    type VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    reject_notes VARCHAR(255),
    verifier_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_requests_users FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_requests_verifiers FOREIGN KEY (verifier_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS user_identities (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    number VARCHAR(30) NOT NULL,
    type VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    expiry_date DATE NOT NULL,
    place_issued VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_identities_users FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
-- Rollback database schema changes

DROP TABLE IF EXISTS user_identities;
DROP TABLE IF EXISTS requests;
DROP TABLE IF EXISTS volunteer_details;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS roles;
