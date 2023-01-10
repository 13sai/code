### read write

```yaml
db:
  separation: true # 配置是否使用读写分离，方便改配置切换
  master: "root:123456@tcp(127.0.0.1:3306)/a0001_chat?charset=utf8mb4&parseTime=True&loc=Local"
  slave:
    - "root:123456@tcp(127.0.0.1:3307)/a0001_chat?charset=utf8mb4&parseTime=True&loc=Local"
    - "root:123456@tcp(127.0.0.1:3308)/a0001_chat?charset=utf8mb4&parseTime=True&loc=Local"
```go
package main

import (
    "fmt"
    "time"

    "github.com/pkg/errors"
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/plugin/dbresolver"
)

var (
    MainDB *gorm.DB
)

func main() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %w \n", err))
    }
    MainDB, err = ConnectDB()
    if err != nil {
        panic(fmt.Errorf("Fatal MainDB config file: %w \n", err))
    }
    err = MainDB.Raw("select version()").Error
    if err != nil {
        logrus.Infof("err=%+v", err)
        return
    }
    logrus.Info("13sai")
}

func ConnectDB() (d *gorm.DB, err error) {
    if viper.GetBool("db.separation") {
        return ConnectRWDB()
    }
    dsn := viper.GetString("db.master")
    d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, errors.Wrap(err, "数据库连接失败")
    }
    db, err := d.DB()
    if err != nil {
        return nil, errors.Wrap(err, "获取数据库实例失败")
    }
    db.SetMaxIdleConns(10)
    db.SetConnMaxLifetime(time.Hour)
    return d, nil
}

func ConnectRWDB() (d *gorm.DB, err error) {
    logrus.Info("使用读写分离")
    dsn := viper.GetString("db.master")
    d, err = gorm.Open(mysql.New(mysql.Config{
        DSN: dsn,
    }))
    if err != nil {
        return nil, err
    }
    replicas := []gorm.Dialector{}
    for i, s := range viper.GetStringSlice("db.slave") {
        cfg := mysql.Config{
            DSN: s,
        }
        logrus.Infof("读写分离-%d-%s", i, s)
        replicas = append(replicas, mysql.New(cfg))
    }

    d.Use(
        dbresolver.Register(dbresolver.Config{
            Sources: []gorm.Dialector{mysql.New(mysql.Config{
                DSN: dsn,
            })},
            Replicas: replicas,
            Policy:   dbresolver.RandomPolicy{},
        }).
            SetMaxIdleConns(10).
            SetConnMaxLifetime(time.Hour).
            SetMaxOpenConns(200),
    )

    return d, nil
}
```