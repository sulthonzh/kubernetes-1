create database if not exists auth;
create database if not exists config;
create database if not exists event;
create database if not exists trace;
create database if not exists router;
grant all privileges on auth.* to 'auth'@'%' identified by 'auth';
grant all privileges on config.* to 'config'@'%' identified by 'config';
grant all privileges on event.* to 'event'@'%' identified by 'event';
grant all privileges on trace.* to 'trace'@'%' identified by 'trace';
grant all privileges on router.* to 'router'@'%' identified by 'router';

