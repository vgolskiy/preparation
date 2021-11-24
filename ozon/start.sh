#!/bin/sh

pg_ctl start
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'facts'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE facts"
psql -c "ALTER USER postgres WITH ENCRYPTED PASSWORD 'ozon';"
psql facts -U postgres -f facts.sql
./mini_webserv
