## 安装

```
git clone https://github.com/wg/wrk.git wrk
cd wrk
make
# move the executable to somewhere in your PATH
sudo cp wrk /somewhere/in/your/PATH
```

默认情况下wrk会使用自带的LuaJIT和OpenSSL，如果你想使用系统已安装的版本，可以使用WITH_LUAJIT和WITH_OPENSSL这两个选项来指定它们的路径。比如：
```
make WITH_LUAJIT=/usr WITH_OPENSSL=/usr
```

## 使用说明

```
wrk <选项> <被测HTTP服务的URL>                            
  Options:                                            
    -c, --connections <N>  跟服务器建立并保持的TCP连接数量  
    -d, --duration    <T>  压测时间           
    -t, --threads     <N>  使用多少个线程进行压测   
                                                      
    -s, --script      <S>  指定Lua脚本路径       
    -H, --header      <H>  为每一个HTTP请求添加HTTP头      
        --latency          在压测结束后，打印延迟统计信息   
        --timeout     <T>  超时时间     
    -v, --version          打印正在使用的wrk的详细版本信息
                                                      
  <N>代表数字参数，支持国际单位 (1k, 1M, 1G)
  <T>代表时间参数，支持时间单位 (2s, 2m, 2h)
```

## 测试

```
wrk -t8 -c200 -d30s --latency  "http://www.13sai.com"

输出：
Running 30s test @ http://www.13sai.com
  8 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    46.67ms  215.38ms   1.67s    95.59%
    Req/Sec     7.91k     1.15k   10.26k    70.77%
  Latency Distribution
     50%    2.93ms
     75%    3.78ms
     90%    4.73ms
     99%    1.35s 
  1790465 requests in 30.01s, 684.08MB read
Requests/sec:  59658.29


Running 30s test @ http://www.13sai.com （压测时间30s）
  8 threads and 200 connections （共8个测试线程，200个连接）
  Thread Stats   Avg      Stdev     Max   +/- Stdev
              （平均值） （标准差）（最大值）（正负一个标准差所占比例）
    Latency    46.67ms  215.38ms   1.67s    95.59%
    （延迟）
    Req/Sec     7.91k     1.15k   10.26k    70.77%
    （处理中的请求数）
  Latency Distribution （延迟分布）
     50%    2.93ms
     75%    3.78ms
     90%    4.73ms
     99%    1.35s （99分位的延迟）
  1790465 requests in 30.01s, 684.08MB read （30.01秒内共处理完成了1790465个请求，读取了684.08MB数据）
Requests/sec:  59658.29 （平均每秒处理完成59658.29个请求）
Transfer/sec:     22.79MB （平均每秒读取数据22.79MB）
```

## lua脚本使用

```
测试：
```
wrk -t500 -c1000 -d1s  --script=data.lua --latency https://github.13sai.com/
```

感觉比ab好用，推荐一下。

