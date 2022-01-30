CREATE TYPE article_states AS ENUM('accepted', 'review', 'rejected');

ALTER TABLE articles ADD COLUMN article_state article_states NOT NULL DEFAULT 'review';
UPDATE articles SET article_state = 'accepted' WHERE accepted = TRUE;
ALTER TABLE articles DROP COLUMN accepted;