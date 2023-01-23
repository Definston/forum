CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL CHECK(LENGTH(email) <= 50),
    nick VARCHAR(32) UNIQUE NOT NULL CHECK(LENGTH(nick) <= 32),
    pass VARCHAR(50) NOT NULL CHECK(LENGTH(email) <= 50));

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INTEGER NOT NULL,
    user_nick NVARCHAR(32) NOT NULL,
    parent_id INT NOT NULL DEFAULT 0,
    content TEXT NOT NULL,
    link TEXT,
    resiever INT,
    FOREIGN KEY (user_id) REFERENCES users(id));

CREATE TABLE IF NOT EXISTS votes (
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    vote BOOLEAN NOT NULL,
    UNIQUE (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id));

CREATE TABLE IF NOT EXISTS tags (
    post_id INTEGER NOT NULL,
    tag TEXT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    UNIQUE (post_id, tag));

PRAGMA foreign_keys = ON;