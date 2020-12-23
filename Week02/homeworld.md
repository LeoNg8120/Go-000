#### Week02 作业题目：

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


#### 说明
以上作业，要求提交到 Github 上面，Week02 作业提交地址：
https://github.com/Go-000/Go-000/issues/8  

评语:
在 dao 层建议把源信息带上，即使不在 wrap 里面，在format 里面也可以展示，慎防底层把错误提示信息修改掉。


老师参考：  
主要是不要在最底层吞掉异常，会有以下两类问题1.直接吞掉异常，替换成自定义异常，没有wrap，2.wrap了自定义异常，没有把源异常打印
```
dao: 

 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))


biz:

if errors.Is(err, code.NotFound} {

}

```

