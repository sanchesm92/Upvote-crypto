-- Active: 1660844667914@@127.0.0.1@5432@upvote
DROP DATABASE IF EXISTS upvote;

CREATE DATABASE upvote;
CREATE TABLE currencies(
    id serial NOT NULL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    vote INT NOT NULL
)

