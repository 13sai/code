偶尔会有代码提交错误，又提交到远程分支了，怎么回退呢？

### 强制回退

```git
# 查看日志，找到对应的commit id
git log
git reset --hard 回退的版本id
git push -f origin 分支名
```

这样回退是清除了回退的版本id之后的提交，连日志都没有了。

### 回退版本

```git
# 查看日志，找到对应的commit id
git log
git reset --soft 回退的版本id
git commit 
```

这样回退是改回回退的版本id之后的提交，日志仍然存在。





