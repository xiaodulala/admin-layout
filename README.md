# admin-layout(后端管理系统脚手架)


# 功能特性

自带用户，菜单 角色 部门 岗位等管理功能。权限包含功能权限和数据权限。


# 快速开始

* 编译

    ```bash
    make build
    ```
  
* 配置

配置confg/appctl.yaml已经application.yaml中的数据库配置

* 初始化数据库

    ```bash
    ./_output/platforms/darwin/amd64/appctl migrate
    ```


* 启动后台管理项目
  
    ```bash
    ./_output/platforms/darwin/amd64/application 
    ```


# 使用指南

* 编译项目,提供makefile项目管理
```bash
# 查看帮助
make help

Usage: make <TARGETS> <OPTIONS> ...

Targets:
  clean             删除编译时产生的所有文件.
  build             根据当前平台编译对应的application可执行文件.
  build.multiarch   编译多平台的可执行文件. See option PLATFORMS.
  lint              代码静态检查.
  install           安装开发项目时的所有工具(网络不好请重试).
  swagger           生成swagger文档.
  serve-swagger     启动swagger接口服务
  tidy              安装go mod 依赖
  image             以当前平台制作镜像
  image.multiarch   生成多平台镜像. See option PLATFORMS.
  help              Show this help info.

```

* 编译出来的命令一共有两个
  * appctl  客户端命令行工具。
  * application 后端应用程序
    
可以分别通过 -h参数查看命令支持的选项参数和子命令


