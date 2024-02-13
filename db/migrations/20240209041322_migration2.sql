-- migrate:up
ALTER TABLE users DROP COLUMN refreshtoken;
CREATE TABLE refreshtoken(UserID INT64 NOT NULL REFERENCES users (ID), RefreshToken string NOT NULL);

-- migrate:down
ALTER TABLE users ADD COLUMN refreshtoken string;
DROP TABLE IF EXISTS refreshtoken;
