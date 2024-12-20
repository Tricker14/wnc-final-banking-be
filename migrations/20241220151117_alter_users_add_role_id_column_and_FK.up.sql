-- Step 1: Add the role_id column after the name column
ALTER TABLE users
ADD COLUMN role_id INT AFTER name;

ALTER TABLE users
ADD CONSTRAINT fk_role_id
FOREIGN KEY (role_id)
REFERENCES roles(id)
ON DELETE SET NULL
ON UPDATE CASCADE;