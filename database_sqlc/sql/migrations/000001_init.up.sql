CREATE TABLE categories (
    ID VARCHAR(36) NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE courses (
    id varchar(36) not null primary key,
    category_id varchar(36) not null,
    name text not null,
    description text,
    price decimal(10,2) not null,
    FOREIGN key (category_id) REFERENCES categories(id)
);