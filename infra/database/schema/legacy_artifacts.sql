CREATE TABLE Users (
    id varchar(36) NOT NULL PRIMARY KEY,
    email varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (id)
);