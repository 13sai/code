## 查看 nginx 日志 qps

```shell
cat /var/log/nginx/access.log | awk '{print $4}' | uniq -c | sort -r | head
cat /var/log/nginx/access.log|awk '{a[$4]+=1;}END{for (i in a) { printf("%s\t%i\n",i,a[i])}}'|sort -gr -k2|head
```

### nohup

```shell
nohup python3 sai.py >/dev/null 2>&1 &
```

### 统计 gz 行数

```shell
zcat ./b2b_4_3_0.csv.gz | sed -n '$='
zcat ./b2b_4_3_0.csv.gz | awk  'END{print NR}'
zcat ./b2b_4_3_0.csv.gz  | grep -n "" |awk -F : 'END{print $1}'
```

### 杀死所有 fpm

```shell
ps -ef|grep php-fpm| grep -v grep |awk '{print $2}'|xargs kill -9
```