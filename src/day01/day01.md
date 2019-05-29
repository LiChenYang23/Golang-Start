# 1, for _,number = range numbers()
    _是占位符，如果不需要使用 foreach中的index的代替

# 2,
    查找map里的键时，有两个选择：要么赋值给一个变量，要么为了精确查找，赋值给两个变量。
    赋值给两个变量时第一个值和赋值给一个变量时的值一样，是map查找的结果值。如果指定了第二个值，就会返回一个布尔标志，来表示查找的键是否存在于map里。
    如果这个键不存在，map会返回其值类型的零值作为返回值，如果这个键存在，map会返回键所对应值的副本。
    ```
    matcher, exists := matchers[feed.Type]
    if !exists {
       matcher = matchers["default"]
    }
    ```

# 3,feeds, err := RetrieveFeeds()
    函数标准返回，返回对象 + 错误
