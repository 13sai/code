## create

```sql
CREATE TABLE `sai` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int unsigned NOT NULL,
  `trend` json DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

## alter

```sql
alter table sai add remark varchar(32) AFTER `trend`;
```

online ddl:

```sql
ALTER table sai add index idx_member_id(member_id), ALGORITHM=INPLACE, LOCK=NONE;
```

## drop

```sql
drop table sai;
```

## truncate

```sql
TRUNCATE table sai;
```

