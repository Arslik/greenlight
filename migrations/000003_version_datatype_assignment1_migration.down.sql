ALTER TABLE movies DROP COLUMN version;
ALTER TABLE movies ADD version int NOT NULL;