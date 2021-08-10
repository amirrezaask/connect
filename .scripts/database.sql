-- DROP old ones :))
DROP TABLE IF EXISTS channel_users;
DROP TABLE IF EXISTS hub_users;
DROP TABLE IF EXISTS hub_permissions;
DROP TABLE IF EXISTS channel_permissions;

DROP TABLE IF EXISTS channels;
DROP TABLE IF EXISTS hubs;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS channel_type_t;

-- Users --
CREATE TABLE IF NOT EXISTS users (
    id varchar(255) primary key,
    username varchar(255),
    email varchar(255),
    password_hash varchar(255)
);
/* CREATE INDEX users_username_idx users(username); */
/* CREATE INDEX users_email_idx users(email); */

-- Hubs

CREATE TABLE IF NOT EXISTS hubs (
    id varchar(255) primary key,
    name varchar(255),
    creator varchar(255) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS hub_users (
    user_id varchar(255) REFERENCES users(id),
    hub_id varchar(255) REFERENCES hubs(id),
    primary key (user_id, hub_id)
);
-- Channels
CREATE TYPE channel_type_t AS ENUM ('voice', 'text');

CREATE TABLE IF NOT EXISTS channels (
    id varchar(255) primary key,
    name varchar(255),
    "type" channel_type_t,
    hub_id varchar(255) REFERENCES hubs(id) NOT NULL
);
CREATE TABLE IF NOT EXISTS channel_users (
    user_id varchar(255) REFERENCES users(id),
    channel_id varchar(255) REFERENCES channels(id),
    primary key(user_id, channel_id)
);
CREATE TABLE IF NOT EXISTS hub_permissions (
    user_id varchar(255) REFERENCES users(id),
    hub_id varchar(255) REFERENCES hubs(id),
    role_name varchar(255),
    primary key(user_id, hub_id, role_name)
);
CREATE TABLE IF NOT EXISTS channel_permissions (
    user_id varchar(255) REFERENCES users(id),
    channel_id varchar(255) REFERENCES channels(id),
    role_name varchar(255),
    primary key(user_id, channel_id, role_name)
);

CREATE TABLE IF NOT EXISTS messages(
    id varchar(255) primary key,
    user_id varchar(255) REFERENCES users(id),
    channel_id varchar(255) REFERENCES channels(id),
    hub_id varchar(255) REFERENCES hubs(id),
    payload varchar(255)
);



