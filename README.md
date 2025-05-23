# Task Tracker

这是一个基于Go语言开发的命令行任务追踪工具，支持任务的增删改查、状态管理及分类列表查看功能，通过JSON文件持久化存储任务数据。

## 功能特性
- 任务基本操作：添加、更新、删除任务
- 状态管理：标记任务为待办（pending）、进行中（in progress）、已完成（done）
- 分类查看：支持查看所有任务、已完成任务、进行中任务、未完成任务
- 数据持久化：自动通过`tasks.json`文件存储任务数据

## 安装与运行
1. 确保已安装Go环境（版本≥1.21）
2. 克隆/下载项目到本地（当前目录结构包含`main.go`、`go.mod`）
3. 运行命令：
   ```bash
   go run main.go [命令] [参数]
   ```

## 命令说明
### 1. 添加任务（add）
- 格式：`task-tracker add <标题> <描述>`
- 示例：`go run main.go add "学习Go语言" "完成基础语法章节"`
- 说明：自动生成递增任务ID，初始状态为`pending`

### 2. 更新任务（update）
- 格式：`task-tracker update <任务ID> <新标题> <新描述>`
- 示例：`go run main.go update 1 "学习Go语言进阶" "完成并发编程章节"`
- 说明：通过任务ID定位需要更新的任务

### 3. 删除任务（delete）
- 格式：`task-tracker delete <任务ID>`
- 示例：`go run main.go delete 1`
- 说明：删除指定ID的任务（不可恢复）

### 4. 标记状态（mark）
- 格式：`task-tracker mark <任务ID> <状态>`
- 状态可选值：`pending`（待办）、`in progress`（进行中）、`done`（已完成）
- 示例：`go run main.go mark 1 "in progress"`
- 说明：修改指定任务的状态

### 5. 查看任务（list）
- 基础格式：`task-tracker list`（查看所有任务）
- 过滤格式：
  - `task-tracker list done`（查看已完成任务）
  - `task-tracker list in-progress`（查看进行中任务）
  - `task-tracker list not-done`（查看未完成任务，即非done状态）
- 示例：`go run main.go list done`

## 数据存储
任务数据自动存储在同级目录的`tasks.json`文件中，格式如下：
```json
[
  {
    "id": 1,
    "title": "学习JSON",
    "description": "掌握JSON序列化协议",
    "status": "pending"
  },
  {
    "id": 2,
    "title": "学习Golang",
    "description": "完成Go语言从入门到精通",
    "status": "in progress"
  }
]
```

## 注意事项
- 首次运行会自动创建空的`tasks.json`文件
- 所有操作（增删改状态）都会自动保存到JSON文件
- 无效参数（如非数字ID、非法状态）会提示错误信息

