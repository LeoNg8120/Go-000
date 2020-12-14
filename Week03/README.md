#### 学习笔记

#### 作业

课程作业布置【作业提交截止时间12月9日（周三）23:59前】

👉Week03 作业题目：
1.基于 errgroup 实现多个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。  
﻿
⚠️以上作业，要求提交到Github上面，Week03作业提交地址：
https://github.com/Go-000/Go-000/issues/69

👉请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。  
﻿
⚠️Github使用教程：https://u.geekbang.org/lesson/51?article=294701

🎈学号查询方式：
PC端登录time.geekbang.org,点击右上角头像进入【我的教室】，左侧头像下方G开头的为学号。

#### comment老师建议
1. 建议使用 errgroup 控制整个流程
2. 除了这个信号量之外，如果其中一个 server 启动的时候，报错也需要考虑。

