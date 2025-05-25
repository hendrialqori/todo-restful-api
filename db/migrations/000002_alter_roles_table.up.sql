START TRANSACTION;

ALTER TABLE roles DROP INDEX unique_name;

ALTER TABLE roles ADD UNIQUE unique_name (name);

COMMIT;