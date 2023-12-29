BEGIN;

ALTER TABLE employee ADD CONSTRAINT unique_position_id UNIQUE (position_id);

ALTER TABLE employee DROP CONSTRAINT fk_position;

ALTER TABLE position
ADD CONSTRAINT fk_position FOREIGN KEY (position_id) REFERENCES employee(position_id);

COMMIT;