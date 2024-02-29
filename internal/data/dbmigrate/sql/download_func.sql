-- Create a table to maintain last 3 inserted downloads
CREATE TABLE downloads_history (
   seq_no INTEGER,
   id UUID,
   source_id UUID,
   attempted_at TIMESTAMP  WITH TIME ZONE NOT NULL,
   downloaded_at TIMESTAMP  WITH TIME ZONE NOT NULL,
   status_code INTEGER,
   headers jsonb,
   body TEXT,

   PRIMARY KEY (id)
);

-- Create a sequence
CREATE SEQUENCE downloads_seq;

-- Create a trigger function
CREATE OR REPLACE FUNCTION maintain_last_three_downloads()
RETURNS TRIGGER AS $$
BEGIN
  -- Insert the new row into the history table
INSERT INTO downloads_history (id, source_id, attempted_at, downloaded_at, status_code, headers, body)
VALUES (nextval(downloads_seq), NEW.id, NEW.source_id, NEW.attempted_at, NEW.downloaded_at, NEW.status_code, NEW.headers, NEW.body);

-- Delete rows older than the last three
DELETE FROM downloads_history
WHERE id NOT IN (
    SELECT id
    FROM downloads_history
    ORDER BY downloaded_at DESC
    LIMIT 3
    );

RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Create a trigger
CREATE TRIGGER maintain_last_three_downloads_trigger
    AFTER INSERT ON downloads
    FOR EACH ROW EXECUTE FUNCTION maintain_last_three_downloads();