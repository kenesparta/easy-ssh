-- Create SSH params table
CREATE TABLE IF NOT EXISTS hosts
(
    "id"            INTEGER PRIMARY KEY,
    "host"          VARCHAR(50) UNIQUE NOT NULL,
    "hostname"      VARCHAR(50)        NOT NULL,
    "user"          VARCHAR(50)        NOT NULL,
    "identities"    BOOLEAN,
    "identity_file" VARCHAR(250),
    "port"          INT DEFAULT 22
);