# Megrez 项目概览

## 项目简介

Megrez（天权星）是一个基于 Go 语言的博客系统，支持跨平台一键部署和自定义主题。项目采用前后端分离架构，后端使用 Gin + GORM 构建 RESTful API 和模板渲染页面，前端管理界面使用 Vue 2.6 + TDesign 构建并通过 `embed.FS` 嵌入到最终二进制文件中，实现单文件部署。

## 技术栈

| 层面 | 技术 |
|------|------|
| Web 框架 | Gin |
| ORM | GORM（SQLite / MySQL） |
| 模板引擎 | Pongo2 |
| 日志 | Zap |
| 缓存 | go-cache（内存缓存） |
| 认证 | golang-jwt |
| Markdown | Lute |
| API 文档 | Swagger / OpenAPI 2.0 |
| 对象存储 | 腾讯云 COS（可选） |
| 前端管理 | Vue 2.6 + TDesign Vue + Vditor |
| Go 版本 | 1.19 |

## 目录结构

```
megrez/
├── main.go                          # 程序入口，Swagger 注解
├── go.mod / go.sum                  # Go 模块定义与依赖校验
├── Makefile                         # 构建自动化（编译、Docker、前端）
├── Dockerfile                       # 多阶段 Docker 镜像构建
├── .gitmodules                      # Git 子模块（默认主题）
│
├── app/                             # 应用层
│   └── app.go                       # Megrez 结构体，初始化与启动流程
│
├── pkg/                             # 核心业务包
│   ├── api/                         # API 处理层
│   │   ├── admin/                   # 管理端 API（需 JWT 认证）
│   │   │   ├── article.go           #   文章 CRUD
│   │   │   ├── category.go          #   分类管理
│   │   │   ├── tag.go               #   标签管理
│   │   │   ├── comment.go           #   评论管理
│   │   │   ├── journal.go           #   日志管理
│   │   │   ├── link.go              #   友链管理
│   │   │   ├── page.go              #   页面管理
│   │   │   ├── attachment.go        #   附件上传
│   │   │   ├── theme.go             #   主题管理
│   │   │   ├── option.go            #   选项设置
│   │   │   ├── setting.go           #   系统设置
│   │   │   ├── install.go           #   安装向导
│   │   │   └── user.go              #   用户登录
│   │   └── openapi/                 # 开放 API（公开访问）
│   │       └── comment.go           #   评论创建与列表
│   │
│   ├── router/                      # 路由层
│   │   ├── router.go                # Gin 引擎创建、中间件注册、路由挂载
│   │   ├── api/                     # API 路由注册
│   │   │   ├── admin.go             #   /api/admin/* 路由组
│   │   │   └── openapi.go           #   /api/* 开放路由组
│   │   └── view/                    # 视图路由（模板渲染页面）
│   │       ├── view.go              #   视图路由注册
│   │       ├── article.go           #   首页、文章详情
│   │       ├── category.go          #   分类页
│   │       ├── journal.go           #   日志页
│   │       ├── page.go              #   自定义页面
│   │       ├── link.go              #   友链页
│   │       ├── search.go            #   搜索页
│   │       ├── about.go             #   关于页
│   │       └── global.go            #   全局模板数据
│   │
│   ├── model/                       # 数据模型层（GORM）
│   │   ├── db.go                    # 数据库连接与迁移
│   │   ├── user.go                  # 用户模型
│   │   ├── article.go               # 文章模型
│   │   ├── category.go              # 分类模型
│   │   ├── tag.go                   # 标签模型
│   │   ├── comment.go               # 评论模型（树形结构）
│   │   ├── journal.go               # 日志模型
│   │   ├── page.go                  # 页面模型
│   │   ├── link.go                  # 友链模型
│   │   ├── menu.go                  # 菜单模型
│   │   ├── option.go                # 系统选项（键值对）
│   │   ├── attachment.go            # 附件模型
│   │   ├── articletag.go            # 文章-标签关联
│   │   ├── articlecategory.go       # 文章-分类关联
│   │   └── themeoption.go           # 主题选项
│   │
│   ├── middleware/                   # HTTP 中间件
│   │   ├── jwt/                     # JWT 认证
│   │   ├── cros/                    # CORS 跨域
│   │   ├── checkinstall/            # 安装状态检查
│   │   └── pongo2gin/               # Pongo2 模板渲染适配
│   │
│   ├── entity/                      # 数据传输实体
│   │   ├── dto/                     # Data Transfer Object
│   │   │   ├── admin/               #   管理端请求/响应
│   │   │   ├── openapi/             #   开放 API 请求/响应
│   │   │   └── pagination.go        #   分页参数
│   │   └── vo/                      # View Object（视图渲染数据）
│   │       ├── article.go
│   │       ├── category.go
│   │       ├── comment.go
│   │       ├── journal.go
│   │       ├── page.go
│   │       ├── tag.go
│   │       ├── link.go
│   │       ├── menu.go
│   │       ├── pagination.go
│   │       └── global.go
│   │
│   ├── config/                      # 配置管理
│   │   ├── app.go                   # 应用配置结构（数据库、调试模式）
│   │   ├── themeconfig.go           # 主题配置
│   │   └── themeinfo.go             # 主题元信息
│   │
│   ├── cache/                       # 缓存
│   │   └── cache.go                 # 基于 go-cache 的内存缓存
│   │
│   ├── log/                         # 日志
│   │   └── log.go                   # Zap 日志封装
│   │
│   ├── bolt/                        # BoltDB（预留，暂未启用）
│   │   └── bolt.go
│   │
│   └── utils/                       # 工具函数
│       ├── jwt/                     # JWT 生成与解析
│       ├── errmsg/                  # 错误码与消息定义
│       ├── dir/                     # 目录操作
│       ├── file/                    # 文件操作
│       ├── uploader/                # 文件上传（本地 / 腾讯云 COS）
│       ├── useragent/               # User-Agent 解析
│       └── zip/                     # ZIP 压缩/解压
│
├── admin/                           # 前端管理界面（Vue 2.6 项目）
│   ├── src/                         # Vue 源代码
│   ├── public/                      # 静态资源
│   ├── dist/                        # 构建产物
│   └── package.json                 # Node.js 依赖
│
├── assets/                          # 嵌入式静态资源
│   ├── admin/                       # 管理界面（embed.FS）
│   │   └── admin.go                 # //go:embed 嵌入前端构建产物
│   └── themes/                      # 主题资源（embed.FS）
│       ├── themes.go                # //go:embed 嵌入默认主题
│       └── default/                 # 默认主题（git submodule）
│
├── docs/                            # API 文档
│   ├── docs.go                      # Swagger 文档注册
│   ├── swagger.yaml                 # OpenAPI 2.0 定义
│   ├── swagger.json                 # JSON 格式
│   └── images/                      # 文档图片 / 效果预览
│
└── build/                           # 编译产物输出目录（gitignore）
```

## 架构概览

### 分层架构

```
┌─────────────────────────────────────────────────────┐
│                    main.go (入口)                     │
├─────────────────────────────────────────────────────┤
│                    app (应用层)                       │
│         初始化：目录 → 日志 → 配置 → DB → 缓存 → 路由  │
├────────────────────────┬────────────────────────────┤
│     router/view        │       router/api            │
│   (模板渲染页面)        │    (RESTful API)            │
├────────────────────────┴────────────────────────────┤
│                  middleware                           │
│          JWT · CORS · 安装检查 · Pongo2               │
├─────────────────────────────────────────────────────┤
│                   api (处理层)                        │
│            admin (管理端) / openapi (开放)             │
├─────────────────────────────────────────────────────┤
│               entity (DTO / VO)                      │
├─────────────────────────────────────────────────────┤
│                 model (数据层)                        │
│            GORM · SQLite / MySQL                     │
├─────────────────────────────────────────────────────┤
│          utils · config · cache · log                │
└─────────────────────────────────────────────────────┘
```

### 启动流程

1. `main.go` 调用 `app.New()` 创建 `Megrez` 实例
2. `Init()` 依次执行以下初始化步骤：
   - `initDir` — 创建或获取 Megrez 主目录（`~/.megrez`）
   - `initLogger` — 初始化 Zap 结构化日志
   - `initConfig` — 从 `config.yaml` 读取配置，失败则使用默认配置（SQLite）
   - `initDAO` — 根据配置连接 SQLite 或 MySQL，执行 AutoMigrate
   - `initCache` — 初始化 go-cache 内存缓存（默认 10 分钟过期）
   - `initRouter` — 创建 Gin 引擎，注册中间件和路由
3. `Run()` 启动 HTTP 服务器（默认端口 8080）

### 请求处理流程

```
客户端请求
    │
    ▼
  Gin Engine
    │
    ├── CORS 中间件
    ├── 安装检查中间件（未安装则重定向到安装页）
    ├── Zap 日志中间件
    └── Recovery 中间件
         │
         ├── /admin/*        → 嵌入式前端管理 SPA
         ├── /api/admin/*    → JWT 认证 → Admin API Handler
         ├── /api/*          → Open API Handler
         ├── /swagger/*      → Swagger UI
         ├── /upload/*       → 静态文件服务
         ├── /themes/*       → 主题静态资源
         └── /*              → 视图路由 → Pongo2 模板渲染
```

## 数据模型

所有模型通过 GORM AutoMigrate 自动建表和迁移：

| 模型 | 说明 | 核心字段 |
|------|------|----------|
| User | 用户 | Username, Nickname, Email, Password, Avatar |
| Article | 文章 | Title, Slug, Content, Summary, Cover, Status |
| Category | 分类 | Name, Slug, Description |
| Tag | 标签 | Name, Slug, Description |
| Comment | 评论 | Content, Author, Email, Type, RootID, ParentID |
| Journal | 日志 | Content, Images, Private, Visits, Likes |
| Page | 页面 | Name, Slug, Content, Type (内置/自定义) |
| Link | 友链 | Name, URL, Logo, Priority |
| Menu | 菜单 | Name, Slug, PageID, Priority |
| Option | 系统选项 | Key, Value（键值对存储） |
| Attachment | 附件 | URL, ThumbURL, FileName, Type |
| ArticleTag | 文章-标签关联 | ArticleID, TagID |
| ArticleCategory | 文章-分类关联 | ArticleID, CategoryID |
| ThemeOption | 主题选项 | ThemeName, Key, Value |

## API 结构

### 管理端 API（`/api/admin/*`）

需要 JWT 认证（登录和安装接口除外）：

- `POST /api/admin/login` — 登录
- `POST /api/admin/install` — 安装向导
- 文章：CRUD `/api/admin/articles`
- 分类：CRUD `/api/admin/categories`
- 标签：CRUD `/api/admin/tags`
- 评论：CRUD `/api/admin/comments`
- 日志：CRUD `/api/admin/journals`
- 友链：CRUD `/api/admin/links`
- 页面：CRUD `/api/admin/pages`
- 附件：CRUD `/api/admin/attachments`
- 主题：`/api/admin/themes`
- 选项：`/api/admin/options`
- 设置：`/api/admin/settings`

### 开放 API（`/api/*`）

公开访问，无需认证：

- `GET /api/{type}/{id}/comments` — 获取评论列表
- `POST /api/comment` — 创建评论

### 视图路由

通过 Pongo2 模板引擎渲染前端页面：

- `/` — 首页（文章列表）
- `/article/:slug` — 文章详情
- `/category/:slug` — 分类文章列表
- `/journal` — 日志页
- `/link` — 友链页
- `/about` — 关于页
- `/search` — 搜索页
- `/:slug` — 自定义页面

## 主题系统

- 主题基于 Pongo2 模板引擎（Django 风格语法）
- 默认主题以 Git submodule 方式管理，位于 `assets/themes/default/`
- 编译时通过 `//go:embed` 将默认主题嵌入二进制文件
- 运行时主题文件释放到 `~/.megrez/themes/` 目录
- 支持通过管理后台上传和切换主题
- 每个主题包含 `theme.yaml`（元信息）和 `config.yaml`（可配置选项）

## 配置

应用配置文件为 `~/.megrez/config.yaml`，首次运行若文件不存在则使用默认配置（SQLite）：

```yaml
database:
  sqlite:
    path: /path/to/megrez.db   # SQLite 数据库路径
  mysql:                        # 或使用 MySQL
    host: localhost
    port: "3306"
    user: root
    password: ""
    name: megrez
debug: true                     # 调试模式
```
