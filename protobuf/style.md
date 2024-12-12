    文件(Files)
        文件名使用小写下划线的命名风格，例如 lower_snake_case.proto
        每行不超过 80 字符
        使用 2 个空格缩进

    包(Packages)
        包名应该和目录结构对应，例如文件在my/package/目录下，包名应为 my.package

    消息和字段(Messages & Fields)
        消息名使用首字母大写驼峰风格(CamelCase)，例如message StudentRequest { ... }
        字段名使用小写下划线的风格，例如 string status_code = 1
        枚举类型，枚举名使用首字母大写驼峰风格，例如 enum FooBar，枚举值使用全大写下划线隔开的风格(CAPITALS_WITH_UNDERSCORES )，例如 FOO_DEFAULT=1
    服务(Services)
        RPC 服务名和方法名，均使用首字母大写驼峰风格，例如service FooService{ rpc GetSomething() }
