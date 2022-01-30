DROP TYPE article_states;

ALTER TABLE articles ADD COLUMN accepted BOOLEAN DEFAULT FALSE;
UPDATE articles SET accepted = TRUE WHERE article_state = 'accepted';
ALTER TABLE articles DROP COLUMN article_state;