FROM postgres:14.1-alpine

ADD upvote.sql /docker-entrypoint-initdb.d/