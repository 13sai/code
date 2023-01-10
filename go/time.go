package main

import (
    "time"
    "fmt"
)

func main() {

    a := time.Now().Unix()
    fmt.Println("时间戳---", a)

    // 2006-01-02 15:04:05 记住这一刻
    b := time.Now().Format("2006-01-02 15:04:05")
    fmt.Println("格式化时间", b)

    /**
    func (t Time) Add(d Duration) Time
    Duration如下
    const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
    )
     */
    c := time.Now().Add(time.Minute * 3)
    fmt.Println("3分钟后时间", c.Format("2006-01-02 15:04:05"))

    /**
        func (t Time) AddDate(years int, months int, days int) Time
    */
    d := time.Now().AddDate(-1, 1,10)
    fmt.Println("时间", d.Format("2006-01-02 15:04:05"))

    // 返回年月日三个值
    fmt.Println(time.Now().Date())
    // 返回时分秒三个值
    fmt.Println(time.Now().Clock())


    fmt.Println(time.Now().Year(), time.Now().Month(), time.Now().Day())
    fmt.Println(time.Now().Weekday(), time.Now().Hour())
    fmt.Println(time.Now().YearDay())

    fmt.Println(time.Since(d))


    // tring返回采用如下格式字符串的格式化时间。
    // "2006-01-02 15:04:05.999999999 -0700 MST"
    fmt.Println(time.Now().String())

    time.AfterFunc(2*time.Second, func() {
        fmt.Println("hello 2s")
    })

    loc, _ := time.LoadLocation("Asia/Shanghai")
    const longForm = "Jan 2, 2006 at 3:04pm (MST)"
    const shortForm = "2006-Jan-02"
    t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
    fmt.Println(t)

    /**
    func ParseInLocation(layout, value string, loc *Location) (Time, error)
     */
    t, _ = time.ParseInLocation(shortForm, "2022-Jul-09", loc)
    fmt.Println(t)

    /**
    func Parse(layout, value string) (Time, error)

    解析一个格式化的时间字符串并返回它代表的时间
    ParseInLocation类似Parse但有两个重要的不同之处。
    第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc；
    第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc
    */
    t, _ = time.Parse(longForm, "Feb 3, 2023 at 7:54pm (PST)")
    fmt.Println(t)


    t, _ = time.Parse(shortForm, "2020-Feb-03")
    fmt.Println(t)

    ch := make(chan int)
    timeout := time.After(time.Second * 2)
    timer := time.NewTimer(time.Second * 4) 
    var i int
    go func() {
        for {
            // i++
            select {
                case <- ch:
                    fmt.Println("channel close")
                    return
                case <- timer.C:
                    fmt.Println("4s的NewTimer定时任务")
                case <- timeout:
                    fmt.Println("4s定时输出")
                case <- time.After(time.Second * 6):
                    fmt.Println("6s到了") 
                // default:
                //     //Sleep 1秒，参数就是上面的Duration
                //     time.Sleep(time.Second * 1)
                //     fmt.Println("go 1s")
            }
        }
    }()
    time.Sleep(time.Second * 15)
    fmt.Println("close----")
    close(ch)
    time.Sleep(time.Second * 2)
}