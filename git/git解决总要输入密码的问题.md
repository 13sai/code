之前每次pull和push总要输入密码，繁琐。操作两步，解决问题！

生成公钥

> cd ~/.ssh
> ssh-keygen -t rsa -C "注释"
拷贝公钥到远程主机

> ssh-copy-id git@xxx.com

ok！