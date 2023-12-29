BEGIN;

ALTER TABLE position ALTER COLUMN position_id DROP DEFAULT;

ALTER TABLE employee ALTER COLUMN position_id SET DEFAULT uuid_generate_v4();

COMMIT;