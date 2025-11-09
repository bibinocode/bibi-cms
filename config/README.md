# 配置管理

配置管理模块支持多种配置格式（YAML、TOML、JSON）和环境变量覆盖。

## 功能特性

- ✅ 支持 YAML、TOML、JSON 三种配置格式
- ✅ 环境变量自动覆盖配置
- ✅ 自动设置默认值
- ✅ 全局单例访问
- ✅ 类型安全的配置结构

## 配置加载

### 1. 从 YAML 加载

```go
err := config.LoadConfigFromYaml("config.yaml")
if err != nil {
    log.Fatal(err)
}
```

### 2. 从 TOML 加载

```go
err := config.LoadConfigFromToml("config.toml")
if err != nil {
    log.Fatal(err)
}
```

### 3. 从环境变量加载

```go
err := config.LoadConfigFromEnv()
if err != nil {
    log.Fatal(err)
}
```

### 4. 组合使用（推荐）

```go
// 先加载配置文件
if err := config.LoadConfigFromYaml("config.yaml"); err != nil {
    log.Fatal(err)
}

// 再用环境变量覆盖
if err := config.LoadConfigFromEnv(); err != nil {
    log.Fatal(err)
}
```

## 配置使用

### 访问配置

```go
// 使用全局变量访问
cfg := config.C()

// 应用配置
fmt.Println(cfg.App.Name)
fmt.Println(cfg.App.Port)

// 数据库配置
fmt.Println(cfg.MySQL.Host)
fmt.Println(cfg.MySQL.Database)

// Redis 配置
fmt.Println(cfg.Redis.Host)
fmt.Println(cfg.Redis.Port)

// 日志配置
fmt.Println(cfg.Log.Level)
fmt.Println(cfg.Log.Format)
```

## 配置文件示例

### YAML 格式 (config.yaml)

```yaml
app:
  name: bibi-cms
  host: 0.0.0.0
  port: 8080
  mode: debug

mysql:
  host: 127.0.0.1
  port: 3306
  user: root
  password: your_password
  database: bibi_cms
  charset: utf8mb4

redis:
  host: 127.0.0.1
  port: 6379
  password: ""
  db: 0

log:
  level: info
  format: json
  output: stdout
  path: logs/app.log
```

### TOML 格式 (config.toml)

```toml
[app]
name = "bibi-cms"
host = "0.0.0.0"
port = 8080
mode = "debug"

[mysql]
host = "127.0.0.1"
port = 3306
user = "root"
password = "your_password"
database = "bibi_cms"
charset = "utf8mb4"

[redis]
host = "127.0.0.1"
port = 6379
password = ""
db = 0

[log]
level = "info"
format = "json"
output = "stdout"
path = "logs/app.log"
```

## 环境变量

环境变量会自动覆盖配置文件中的值，支持以下环境变量：

### 应用配置
- `APP_NAME` - 应用名称
- `APP_HOST` - 监听地址
- `APP_PORT` - 监听端口
- `APP_MODE` - 运行模式 (debug/release/test)

### MySQL 配置
- `MYSQL_HOST` - MySQL 主机
- `MYSQL_PORT` - MySQL 端口
- `MYSQL_USER` - MySQL 用户名
- `MYSQL_PASSWORD` - MySQL 密码
- `MYSQL_DATABASE` - 数据库名
- `MYSQL_CHARSET` - 字符集

### Redis 配置
- `REDIS_HOST` - Redis 主机
- `REDIS_PORT` - Redis 端口
- `REDIS_PASSWORD` - Redis 密码
- `REDIS_DB` - Redis 数据库编号

### 日志配置
- `LOG_LEVEL` - 日志级别 (debug/info/warn/error)
- `LOG_FORMAT` - 日志格式 (json/text)
- `LOG_OUTPUT` - 日志输出 (stdout/file)
- `LOG_PATH` - 日志文件路径

## 默认值

如果配置文件和环境变量都未设置，系统会使用以下默认值：

```go
App:
  Name: "bibi-cms"
  Host: "0.0.0.0"
  Port: 8080
  Mode: "debug"

MySQL:
  Host: "127.0.0.1"
  Port: 3306
  Charset: "utf8mb4"

Redis:
  Host: "127.0.0.1"
  Port: 6379
  DB: 0

Log:
  Level: "info"
  Format: "json"
  Output: "stdout"
```

## 使用示例

```go
package main

import (
    "fmt"
    "log"

    "github.com/bibinocode/bibi-cms/config"
)

func main() {
    // 加载配置
    if err := config.LoadConfigFromYaml("config.yaml"); err != nil {
        log.Fatal(err)
    }

    // 环境变量覆盖
    if err := config.LoadConfigFromEnv(); err != nil {
        log.Fatal(err)
    }

    // 使用配置
    cfg := config.C()

    fmt.Printf("Starting %s on %s:%d\n",
        cfg.App.Name,
        cfg.App.Host,
        cfg.App.Port,
    )

    fmt.Printf("MySQL: %s@%s:%d/%s\n",
        cfg.MySQL.User,
        cfg.MySQL.Host,
        cfg.MySQL.Port,
        cfg.MySQL.Database,
    )
}
```