BEGIN;

ALTER TABLE employee DROP COLUMN employee_id;
ALTER TABLE employee ADD COLUMN employee_id BIGINT;

COMMIT;