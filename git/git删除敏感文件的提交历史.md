
不知道你有没有出现过 git 误提交敏感信息的情况，如果是私有库还好，加入是公有库就非常危险了，我之前的应对措施是-----整个库删除，重新建个仓库提交。但实际没必要这样，比较 git history 也是非常重要的。

```sh
# config.yaml 为敏感文件
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch config.yaml' --prune-empty --tag-name-filter cat -- --all
# 随便修改个文件，做点无关痛痒的修改啥的
git commit -am "13sai" 
# 强制 push
git push origin --force --all
```

That's all!
