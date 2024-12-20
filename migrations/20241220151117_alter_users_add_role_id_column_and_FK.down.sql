-- Step 1: Drop the foreign key constraint
ALTER TABLE users
DROP FOREIGN KEY fk_role_id;

-- Step 2: Drop the role_id column
ALTER TABLE users
DROP COLUMN role_id;