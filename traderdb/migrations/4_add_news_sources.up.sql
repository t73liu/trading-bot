CREATE TABLE IF NOT EXISTS news_sources (
  id text PRIMARY KEY,
  name text NOT NULL,
  description text NOT NULL,
  url text NOT NULL
);

COMMENT ON COLUMN news_sources.id IS 'News API unique identifier.';

CREATE TABLE IF NOT EXISTS user_news_sources (
  user_id INTEGER REFERENCES users(id) NOT NULL,
  news_source_id text REFERENCES news_sources(id) NOT NULL,

  PRIMARY KEY (user_id, news_source_id)
);
