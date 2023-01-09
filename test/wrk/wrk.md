### wrk 

[wrk github](https://github.com/wg/wrk.git)

安装
```
git clone https://github.com/wg/wrk.git wrk
cd wrk
make
```

# move the executable to somewhere in your PATH

> sudo cp wrk /somewhere/in/your/PATH

默认情况下 wrk 会使用自带的 LuaJIT 和 OpenSSL，如果你想使用系统已安装的版本，可以使用 WITH_LUAJIT 和 WITH_OPENSSL 这两个选项来指定它们的路径。比如：

> make WITH_LUAJIT=/usr WITH_OPENSSL=/usr



```shell
wrk -t20 -c60 -d60s  --script=data.lua --latency http://127.0.0.1:8888/api
```

```
-c, --connections: total number of HTTP connections to keep open with
                   each thread handling N = connections/threads

-d, --duration:    duration of the test, e.g. 2s, 2m, 2h

-t, --threads:     total number of threads to use

-s, --script:      LuaJIT script, see SCRIPTING

-H, --header:      HTTP header to add to request, e.g. "User-Agent: wrk"

    --latency:     print detailed latency statistics

    --timeout:     record a timeout if a response is not received within
                   this amount of time.
```