BEGIN;

DROP INDEX IF EXISTS idx_firstname;
DROP INDEX IF EXISTS idx_lastname;
DROP INDEX IF EXISTS idx_firstname_lastname;

COMMIT;