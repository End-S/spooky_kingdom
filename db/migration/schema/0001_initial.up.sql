CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE publishers (
    publisher_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, 
    name VARCHAR(255) NOT NULL, 
    label VARCHAR(255) NOT NULL
);

CREATE TYPE article_types AS ENUM('ghost', 'ufo', 'weird');

CREATE TABLE articles (
    article_id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    publisher_id uuid NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    date_published timestamp without time zone NOT NULL,
    date_retrieved timestamp without time zone NOT NULL DEFAULT CURRENT_DATE,
    title VARCHAR(2000) NOT NULL UNIQUE,
    description VARCHAR(500),
    link VARCHAR(2000) NOT NULL,
    article_type article_types,
    FOREIGN KEY (publisher_id)
    REFERENCES publishers (publisher_id)
    ON UPDATE CASCADE ON DELETE CASCADE,
    accepted BOOLEAN DEFAULT FALSE
);
