## internal 目录

私有的应用程序代码库。[Golang 编译器强制执行](https://go.dev\/doc\/go1.4#internalpackages)，不会被他人导入的代码

internal 中的目录结构，多应用时可以具备 
- /internal/app/app1
- /internal/pkg
区分应用程序的部分和应用程序共享的部分

而如果是微服务，则直接 /internal/myapp/ 就够了 

## pkg 目录

外部应用程序可以使用的裤代码，可以被其他项目导入使用。放在此处的代码要慎重。


### 区分理解 [pkg 和 internal](https://travisjeffery.com\/b\/2019\/11\/i-ll-take-pkg-over-internal\/)

当根目录保护大量非 Go 组件和目录时，也应该放到 pkg 中，如果项目很小，非必要使用。主要看非 Go 组件多不多

所以微服务里不用 pkg 的情景比较正常，如果有公用代码，可以直接做成一个库，供其他项目复用。

权限服务中需要用到用户服务的 struct User 时，可以考虑把 struct 放在用户的 pkg 中，供权限服务引用

## vendor 目录

Go module 已经成熟，不需要 vendor 目录来手动管理程序的依赖关系

## api 目录

项目对外提供的 API 文件。如 OpenAPI/Swagger/JSON/protobuf 等定义文件

- /api/protobuf_stark_cas/test/test.pb.go
- /api/protobuf_stark_cas/test/test.proto


## build 目录

打包和 CI 所需的文件

## configs 目录

配置文件

## deployments 或 deploy(k8s) 目录

存放 docker-compose 之类的部署用文件

## scripts 目录

用于执行各种构建、安装、分析等操作的脚本

可以让根目录的 Makefile 来调用它们，使得 Makefile 比较精简

## test 目录

外部测试应用程序和测试数据

## assets 目录

项目中使用的静态资源，如图片

## docs 目录

设计和用户文档，如使用文档，非数据定义文档，那些在 API 目录

## examples 目录

使用示例程序

## tools
此项目的支持工具，可以从 /pkg 和 /internal 导入代码结合使用
