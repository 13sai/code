在服务器上使用 git pull 遇到了一些问题，如果你也遇到，希望我的分享对你有所帮助。

### No such file or directory.

先是遇到这个错误

```shell
git pull
Warning: Identity file /home/sai/.ssh/sai not accessible: No such file or directory.
```

提示是没有文件，我们去新建一下即可。

但是应该新建文件里面是什么内容呢？

查看了下 git config

```shell
git config -l

safe.directory=/src/abc
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true
core.sshcommand=ssh -i ~/.ssh/sai
remote.origin.url=git@github.com:13sai/code.git
```

可以看到定义了 sshcommand

对于 core.sshCommand 
> 如果设置了这个变量，当git fetch和git push需要连接到远程系统时，它们将使用指定的命令而不是ssh。该命令与GIT_SSH_COMMAND环境变量的形式相同，在设置环境变量时将被覆盖。

我们复制 ~/.ssh/id_rsa 即可。

> cp ~/.ssh/id_rsa ~/.ssh/sai

完成这一步之后，有遇到问题。

```shell
git pull

ERROR: Repository not found.
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.
```

这个问题花了我不少时间，去github ssh keys 发现服务器上的 id_rsa.pub 已经加了。

最后重新生成

```shell
ssh-keygen -t rsa -C "13sai"
```

需要一步步确认，会发现 id_rsa.pub 和 id_rsa 已经生成。前往 github 添加 id_rsa.pub 。

```shell
git pull

Already up to date.
```