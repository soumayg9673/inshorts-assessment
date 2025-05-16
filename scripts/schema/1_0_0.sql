-- Create source table
CREATE TABLE IF NOT EXISTS news_sources (
    id SERIAL PRIMARY KEY,
    source_name VARCHAR(150) NOT NULL
);

CREATE EXTENSION postgis;

CREATE INDEX idx_news_sources_source_name ON news_sources (source_name);

-- Create category table
CREATE TABLE IF NOT EXISTS news_categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    category_identifier VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX idx_news_categories_category_identifier ON news_categories (category_identifier);
CREATE INDEX idx_news_categories_category_name ON news_categories (category_name);

-- Create article table
CREATE TABLE IF NOT EXISTS news_articles (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    url TEXT NOT NULL,
    publication_date TIMESTAMP,
    source_id INT,
    relevance_score FLOAT,
    ll_summary TEXT,
    location GEOGRAPHY(Point, 4326),
    FOREIGN KEY (source_id) REFERENCES news_sources(id)
);

-- Publication date filtering (latest news, range queries)
CREATE INDEX idx_news_articles_publication_date ON news_articles (publication_date);

-- Relevance-based sorting/filtering
CREATE INDEX idx_news_articles_relevance_score ON news_articles (relevance_score);

-- Foreign key filtering (e.g., by source)
CREATE INDEX idx_news_articles_source_id ON news_articles (source_id);

-- Geospatial index (MUST for location queries)
CREATE INDEX idx_news_articles_location ON news_articles USING GIST (location);


-- Create article category table
CREATE TABLE IF NOT EXISTS news_article_categories (
    article_id UUID,
    category_id INT,
    PRIMARY KEY (article_id, category_id),
    FOREIGN KEY (article_id) REFERENCES news_articles(id),
    FOREIGN KEY (category_id) REFERENCES news_categories(id)
);

-- Fast lookups by article
CREATE INDEX idx_article_categories_article_id ON news_article_categories (article_id);

-- Fast lookups by category
CREATE INDEX idx_article_categories_category_id ON news_article_categories (category_id);
