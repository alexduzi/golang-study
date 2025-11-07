CREATE TABLE IF NOT EXISTS `categories` (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS `courses` (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    category_id INT not null,
    name text not null,
    description text,
    FOREIGN key (category_id) REFERENCES categories(ID)
);