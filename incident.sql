CREATE DATABASE incident_db;

USE incident_db;

CREATE TABLE incidents(
    id int auto_increment primary key,
    reporter varchar(100) not null,
    description varchar(250) not null,
    status varchar(20) not null DEFAULT 'pendiente',
    created_at timestamp DEFAULT current_timestamp
);