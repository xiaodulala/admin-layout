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
  clean             Remove all files that are created by building.
  build             Build source code for host platform.
  build.multiarch   Build source code for multiple platforms. See option PLATFORMS.
  lint              Check syntax and styling of go sources.
  install           Install iam system with all its components.
  swagger           Generate swagger document.
  serve-swagger     Serve swagger spec and docs.
  tidy
  image             Build docker images for host arch
  image.multiarch   Build docker images for multiple platforms. See option PLATFORMS.
  help              Show this help info.
```

* 编译出来的命令一共有两个
  * appctl  客户端命令行工具。
  * application 后端应用程序
    
可以分别通过 -h参数查看命令支持的选项参数和子命令


