# zkgo
对 [go-zookeeper/zk](https://github.com/samuel/go-zookeeper) 的封装，并利用 [grpc]() 实现服务注册

目录结构
```
zkgo
──────
    │   README.md                   // 说明文件
    │   build                       // 生成二进制文件脚本
    │
    └───cli
    │   │   main.go                 // 主函数文件
    │
    │
    └───zk
    │   │   zkOperator.go           // 对go-zookeeper的封装
    │   
    └───kit
    │   │   
    │   └───endpoints               
    │   │   │   endpoints.go        // grpc 端点层
    │   │   │   middlewares.go      // grpc 中间件（装饰器） 
    │   │   
    │   └───service
    │   │   │   service.go          // grpc 服务层
    │   │   
    │   └───transport
    │       │   tranport.go         // grpc 网络层
    │
    └───utils
    │   │   
    │   └───logger              
    │   │   │   logger.go           // 日志
    │   │     
    │   │   converter.go            // 数据格式转换
    │
    │
    └───pkg  
        │   processer.go            // 服务注册
        
```