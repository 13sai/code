### python遍历list并删除元素

以删除 list 中0 为 🌰

1. 可以使用filter过滤返回新的List

```python
lst = [1, 1, 0, 2, 0, 0, 8, 3, 0, 2, 5, 0, 2, 6]

lst = filter(lambda x: x != 0, lst)
print(lst)
```
这样可以安全删除列表中值为 0 的元素了，filter包括两个参数，分别是function和list，filter把传入的函数依次作用于每个元素，然后根据返回值是True还是False来决定是保留还是丢弃该元素。

2. 列表解析

```python
lst = [x for x in lst if x != 0]
print(lst)
```

3. 或者遍历拷贝的List，操作原始的List

```python
for item in lst[:]:
    if item == 0:
        lst.remove(item)
print(lst)
```

4. 用while循环来搞定，每次循环都先会判断 0 是否在列表中

```python
while 0 in lst:
    lst.remove(0)
print(lst)
```

5. 倒序循环遍历

```python
for item in range(len(lst) - 1, -1, -1):
    if lst[item] == 0:
        del lst[item]

print(lst)
```


### json dump 不带空格

```python
json.dumps({"message": "data is null", "status": 1001}, separators=(',', ':'))
```