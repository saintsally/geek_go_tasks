go进阶训练营作业

第二周作业

/week_2_task

问： 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答： 应该Wrap这个error。记录具体sql操作导致sql.ErrNoRows的日志，方便归因。可以在处理错误的地方加入更详细的堆栈信息，方便定位以及解决问题。

第三周作业

/week_3_task

问： 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

答： 首先实现 HTTP server 的启动和关闭， 监听 linux 的 signal信号，支持 kill -9 或 Ctrl+C 的中断操作操作， errgroup 实现多个 goroutine 的级联退出。