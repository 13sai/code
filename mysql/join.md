### JOIN 的执行顺序

以下是JOIN查询的通用结构
```sql
SELECT <row_list> 
FROM <left_table> 
    <inner|left|right> JOIN <right_table> 
ON <join condition> 
WHERE <where condition>
```

执行顺序(SQL语句里第一个被执行的总是FROM子句):

1. FROM: 对左右两张表执行笛卡尔积，产生第一张表a。行数为n*m（n为左表的行数，m为右表的行数)
2. ON: 根据ON的条件逐行筛选将结果插入新表b中
3. JOIN: 添加外部行，如果指定了LEFT JOIN(LEFT OUTER JOIN)，则先遍历一遍左表的每一行，其中不在b的行会被插入到b，该行的剩余字段将被填充为NULL，形成表c（RIGHT JOIN同理）。但如果指定的是INNER JOIN，就不添加外部行，上述插入过程被忽略
4. WHERE: 对c进行条件过滤，满足条件的行被输出到d
5. SELECT: 取出 d 的指定字段返回

### join

- INNER JOIN ON: 返回左右表互相匹配的所有行（不执行上面的第三步）
- LEFT JOIN ON: 返回左表的所有行，若某些行在右表里没有相对应的匹配行，则将右表的列在新表中置为 NULL
- RIGHT JOIN ON: 返回右表的所有行，若某些行在左表里没有相对应的匹配行，则将左表的列在新表中置为 NULL