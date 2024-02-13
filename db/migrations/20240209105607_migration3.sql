-- migrate:up
CREATE TABLE likes(UserID INT64 NOT NULL REFERENCES users (ID), ArticleID INT64 NOT NULL REFERENCES articles (ID));

-- migrate:down
DROP TABLE IF EXISTS likes;
