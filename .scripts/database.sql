-- Users --
CREATE TABLE IF NOT EXISTS users (
    id varchar(255) primary key,
    username varchar(255),
    email varchar(255),
    password_hash varchar(255)
);
CREATE INDEX users_username_idx (username);
CREATE INDEX users_email_idx (email);

-- Hubs
CREATE TABLE IF NOT EXISTS hubs (
    id varchar(255) primary key,
    name varchar(255),
    creator varchar(255) REFERENCES users(id)
);

-- Channels
CREATE TYPE channel_type_t AS ENUM ('voice', 'text');
CREATE TABLE IF NOT EXISTS channels (
    id varchar(255) primary key,
    name varchar(255),
    'type' channel_type_t
);

CREATE TABLE IF NOT EXISTS hub_users (
    user_id varchar(255) REFERENCES users(id),
    hub_id varchar(255) REFERENCES hubs(id),
    primary key (user_id, hub_id)
);

CREATE TABLE IF NOT EXISTS hub_permissions (
    user_id varchar(255) REFERENCES users(id),
    hub_id varchar(255) REFERENCES hubs(id),
    premission bigint,
    primary key(user_id, channel_id)
);

CREATE TABLE IF NOT EXISTS channel_permissions (
    user_id varchar(255) REFERENCES users(id),
    channel_id varchar(255) REFERENCES channels(id),
    premission bigint,
    primary key(user_id, channel_id)
);

CREATE TABLE IF NOT EXISTS channel_users (
    user_id varchar(255) REFERENCES users(id),
    channel_id varchar(255) REFERENCES channels(id),
    primary key(user_id, channel_id)
);

