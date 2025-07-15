# Sql

## Connect
```bash
mysql -u root -p
mysql -h <host> -P <port> -u <username> -p -D <database>
```

Ctrl+D to quit

## Create
```sql
CREATE TABLE tab_name (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(256),
    age INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_id (id)
);
```


## Insert

```sql
INSERT INTO tab_name
VALUES (val1, val2);

INSERT INTO tab_name (col1, col2)
VALUES (val1, val2);
```

```sql
INSERT INTO tab_name (col1, col2)
VALUES (val1, val2), (val3, val4);

```

If field has type datetime or timestamp, the following will work
```sql
INSERT INTO tab_name VALUES (1, "2020-01-01 00:33:11");

```

## Alter

```sql
ALTER TABLE tab_name ADD COLUMN income INT;
ALTER TABLE tab_name DROP COLUMN name;
ALTER TABLE tab_name RENAME COLUMN name TO primary_name;
-- change column type
ALTER TABLE tab_name MODIFY COLUMN name varchar(255);

ALTER TABLE tab_name RENAME TO new_tab_name;

```

## User

1. create someone named 'username'
after @, it can be 'localhost' or '%': host that's allowed to connect

```sql
CREATE USER 'username'@'%' IDENTIFIED BY 'password';
```

2. grant privileges

- readonly
```sql
GRANT SELECT ON dbname.* TO 'username'@'%';
```

- all priviledges
```sql
GRANT ALL PRIVILEGES ON dbname.* TO 'username'@'%';
```

- other priviledges: `SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX`


3. refresh
```sql;
FLUSH PRIVILEGES
```

4. revoke (if necessary)
```sql
REVOKE ALL PRIVILEGES ON dbname.* FROM 'username'@'%';
```

5. show grants

```sql
SHOW GRANTS FOR 'username'@'%';
```

6. delete user
```sql
DROP USER 'username'@'%';
```


7. check all users
```sql
SELECT User, Host FROM mysql.user;
```


## Data types

| Type        | Size    | Range                     | Desc |
|-------------|---------|---------------------------|------|
| INT/INTEGER | 4 bytes | (-2147483648，2147483647) |      |
| BIGINT      | 8 bytes |                           |      |
| FLOAT       | 4 bytes | (-3.4e38, 3.4e38)         |      |
| DOUBLE      | 8 bytes |                           |      |
|-------------|---------|---------------------------|------|


| Type      | Size    | Range | Desc                                                                    |
|-----------|---------|-------|-------------------------------------------------------------------------|
| DATETIME  | 8 bytes |       | YYYY-MM-DD hh:mm:ss, doesn't change with timezone                       |
| TIMESTAMP | 4 bytes |       | YYYY-MM-DD hh:mm:ss, stored as UTC and conversed to localtime when read |
| DATE      | 3 bytes |       | YYYY-MM-DD                                                              |

Use TIMESTAMP if multi-time zone deployment



| Type    | Size          | Range | Desc                                                                                                               |
|---------|---------------|-------|--------------------------------------------------------------------------------------------------------------------|
| CHAR    | 0-255 bytes   |       | fixed width, if CHAR(10) and row only occupies 3 chars, the 7 left chars will be space padded (max value of n=255) |
| VARCHAR | 0-65535 bytes |       | its space depends on actual chars                                                                                  | 
| TEXT    | 0-65535 bytes |       | for longer text, less efficient, stored out of table, need manual specification for index                          |

CHAR(n), n specified fixed number of characters
VARCHAR(n), n specifies max number of characters.

CHAR: for fixed characters like identity number, phone number, typically faster
VARCHAR: more convenient but slower
 
for both CHAR(n) and VARCHAR(n), if exceed size number n, error will be reported


## Index

```sql
CREATE INDEX index_name ON tab_name (col1, col2);
DROP INDEX index_name ON tab_name;
SHOW INDEX FROM tab_name;
```



## functions

e.g. select lower('ABCD')

LOWER(s)
UPPER(s)
CONCAT(s1,s2,..)
TRIM(s)

ABS(n)
EXP(n)




