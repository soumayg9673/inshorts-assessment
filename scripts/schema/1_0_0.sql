-- Create source table
CREATE TABLE IF NOT EXISTS news_sources (
    id SERIAL PRIMARY KEY,
    source_name TEXT NOT NULL,
    source_identifier TEXT
);

-- Create category table
CREATE TABLE IF NOT EXISTS news_categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    category_identifier VARCHAR(255) NOT NULL
);

-- Create article table
CREATE TABLE IF NOT EXISTS news_articles (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT UNIQUE NOT NULL,
    publication_date TIMESTAMP,
    source_id INT,
    relevance_score FLOAT,
    location GEOGRAPHY(Point, 4326),
    FOREIGN KEY (source_id) REFERENCES news_sources(id)
);

-- Create article category table
CREATE TABLE IF NOT EXISTS news_article_categories (
    article_id UUID,
    category_id INT,
    PRIMARY KEY (article_id, category_id),
    FOREIGN KEY (article_id) REFERENCES news_articles(id),
    FOREIGN KEY (category_id) REFERENCES news_categories(id)
);
