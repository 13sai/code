## insert

```sql
insert into `sai`(member_id, trend, created_at) values(13, '{"a":1}', now());


insert into `sai`(member_id, trend, created_at) 
values(14, '{"a":2}', now()), (56, '{"a":2}', now());
```

## delete

```sql
delete from `sai` where member_id = 56;
```

## update 

```sql
update `sai` set updated_at = now() where member_id = 13;
```

## select

```sql
select * from `sai` limit 2;
```

## cte 

```sql
WITH temp AS (
		SELECT id, member_id
		FROM sai
		ORDER BY id
		LIMIT 10000
)
SELECT a.*, b.type
FROM temp a
	LEFT JOIN oscome b ON a.member_id = b.id
WHERE b.type = 10;
```

## json_table

```sql
select a.id,a.li_vanity,a.full_name,b.* from sai a, 
json_table(a.exp,'$[*]' columns(         
    title varchar(32) path '$.title',         
    company varchar(64) path '$.company_name' ))
b where a.id in (13, 56) 
order by company desc;
```
