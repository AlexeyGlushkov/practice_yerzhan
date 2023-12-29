BEGIN;

ALTER TABLE employee DROP CONSTRAINT unique_position_id;

ALTER TABLE position DROP CONSTRAINT fk_position;

ALTER TABLE position
ADD CONSTRAINT fk_position FOREIGN KEY (position_id) REFERENCES position (position_id);
    
COMMIT;