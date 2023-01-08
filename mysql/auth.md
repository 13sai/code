## user

```sql
CREATE USER '13sai'@'localhost' IDENTIFIED BY '12345678';
CREATE USER '13sai'@'%' IDENTIFIED BY '12345678';
```

## grant

```sql
GRANT ALL PRIVILEGES ON directory . sai TO '13sai'@'localhost';

GRANT ALL PRIVILEGES ON * . * TO '13sai'@'localhost';

GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* TO '13sai'@'localhost' WITH GRANT OPTION;
```

常见的可授权账号使用的权限

- ALL PRIVILEGES ：允许 MySQL 用户完全访问指定的数据库（或者如果没有选择数据库，则可以跨系统进行全局访问）
- CREATE：允许他们创建新表或数据库
- DROP：允许他们删除表或数据库
- DELETE：允许他们从表中删除行
- INSERT：允许他们向表中插入行
- SELECT：允许他们使用 SELECT 命令来读取数据库
- UPDATE：允许他们更新表行
- GRANT OPTION：允许他们授予或删除其他用户的权限

## flush

```sql
FLUSH PRIVILEGES;
```

## revoke 

```sql
REVOKE CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD on *.* FROM '13sai'@'localhost';
REVOKE ALL PRIVILEGES on *.* FROM '13sai'@'localhost';
```

## rename
```sql
RENAME USER 'sai'@'localhost' to '13sai'@'%';
```
## drop

```sql
DROP USER '13sai'@'%';
```

## role

### create

```sql
CREATE ROLE 'app_developer', 'app_read', 'app_write';
```

### grant

```sql
GRANT ALL ON app_db.* TO 'app_developer';
GRANT SELECT ON app_db.* TO 'app_read';
GRANT INSERT, UPDATE, DELETE ON app_db.* TO 'app_write';
```

```sql
CREATE USER 'dev1'@'localhost' IDENTIFIED BY 'dev1pass';
CREATE USER 'read_user1'@'localhost' IDENTIFIED BY 'read_user1pass';
CREATE USER 'read_user2'@'localhost' IDENTIFIED BY 'read_user2pass';
CREATE USER 'rw_user1'@'localhost' IDENTIFIED BY 'rw_user1pass';
```

```sql
GRANT 'app_developer' TO 'dev1'@'localhost';
GRANT 'app_read' TO 'read_user1'@'localhost', 'read_user2'@'localhost';
GRANT 'app_read', 'app_write' TO 'rw_user1'@'localhost';
```