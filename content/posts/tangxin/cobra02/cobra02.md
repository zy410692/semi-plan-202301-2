---
title: "作业: cobra - 02 读取配置配置文件"
categories:
  - 半月刊
tags:
  - 半月刊202301下
  - code
date: "2023-01-13T15:19:07+08:00"
lastmod: "2023-01-13T15:19:07+08:00"
toc: true
draft: false
hiddenFromHomePage: false

#  提交作业修改一下内容
pinTop: true
originAuthor: homework
originLink: ""
---


# 作业: cobra - 02 读取配置配置文件

## 作业要求

1. 使用 https://github.com/spf13/cobra 实现命令工具
2. 命令具有以下参数
    1. `--config` , `-c` 配置文件

**配置文件如下**

```yaml
# config.yml
name: zhangsan
age: 20
```

3. 将配置文件保存为 `JSON` 格式 

```bash
$ cat config.json
```

输出结果

```json
{
    "name":"zhangsan",
    "age": 20
}
```


## 作业解析


`json` 和 `yaml` 是 **最常用的** 配置文件类型， 除此之外还有 `ini, toml, xml` 等。

**解析方法** 一般是
1. `Marshal(v any) ([]byte, error)` 将结构体 **解析** 成 `[]byte` 类型。
2. `Unmarshal(data []byte, v any) error` 将 `[]byte` **映射** 到结构体中。 这里的 `v` 需要是 **指针类型**

另外， **不同的** 解析库对应的的 **方法名称** 和 **实现逻辑** 也不尽相同， 使用的时候需要自己研究。


**常用的解析库**

1. `json`: `encoding/json`
2. `yaml`: `gopkg.in/yaml.v3`, `gopkg.in/yaml.v2` **v2 和 v3** 有区别， 自己研究。


