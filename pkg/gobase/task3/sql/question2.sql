DELIMITER $$

CREATE PROCEDURE transfer_funds(IN from_id INT, IN to_id INT, IN amount DECIMAL(10, 2))
BEGIN
    DECLARE from_balance DECIMAL(10, 2);

    START TRANSACTION;

    -- 查询余额并锁定改行，防止其他事务修改
    SELECT balance INTO from_balance
    FROM accounts
    WHERE id = from_id
    FOR UPDATE;

    -- 判断余额是否足够
    IF from_balance >= amount THEN
       -- 更新余额
       UPDATE accounts SET balance = balance - amount WHERE id = from_id;
       UPDATE accounts SET balance = balance + amount WHERE id = to_id;

       -- 记录交易
       INSERT INTO transactions (from_account_id, to_account_id, amount)
       VALUES (from_id, to_id, amount);

       COMMIT;
    ELSE
       ROLLBACK;
    END IF;
END$$

DELIMITER;