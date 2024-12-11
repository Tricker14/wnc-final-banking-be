CREATE TABLE accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    number CHAR(12) UNIQUE NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TRIGGER create_account_after_customer
    AFTER INSERT ON customers
    FOR EACH ROW
BEGIN
    DECLARE account_number CHAR(12);
    DECLARE is_unique BOOLEAN DEFAULT FALSE;

    WHILE NOT is_unique DO

        SET account_number = LPAD(FLOOR(RAND() * 1000000000000), 12, '0');

        IF NOT EXISTS (
            SELECT 1 FROM accounts WHERE number = account_number
        ) THEN
            SET is_unique = TRUE;
END IF;
END WHILE;

INSERT INTO accounts (customer_id, number, balance)
VALUES (NEW.id, account_number, 0);
END;

