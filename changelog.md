# 版本changelog说明

## 1.0.8
1. [feat] - 增强 dbhelper 对自定义字符串类型的支持并添加 nil 参数检查
2. [feat] - 新增任务规划模式配置
3. [feat] - 新增业务知识网络数据源
4. [opt]  - 环境变量相关优化


## 1.0.10
1. [feat] - 增加x-account-id和x-account-type相关支持
2. [feat] - 配置增加agent和tool工具调用超时时间
3. [feat] - 新增业务域支持功能及资源关联管理（批量关联/解关联/查询）
4. [feat] - 新增获取业务域下所有Agent/AgentTpl ID列表接口
5. [feat] - 新增HTTP请求日志记录器
6. [opt] - DB错误合并优化及InsertBuilder批量插入校验日志增强

## 1.1.0
1. [feat] - 添加业务域关联功能及HTTP请求日志
2. [refact] - 重构业务域服务、Agent导入及已发布代理列表相关代码
3. [chore] - 移除已发布智能体列表缓存及自定义空间相关代码
4. [test] - 添加MockGen生成的mock文件

## 1.1.1
1. [opt] - 移除 CustomSpace 资源类型相关代码

## 1.1.2
1. [feat] - cutil包新增 RemoveKeyFromJSON 函数，支持从 JSON 字符串中删除指定 path 的 key

## 1.1.3
1. [feat] - 新增 Gin 请求日志中间件 (ginrequestlogger) 及 HTTP 请求日志记录优化
2. [feat] - 新增 StatusThreeState 三态状态枚举
3. [feat] - 新增资源类型设置接口 (ResourceTypeSet)
4. [chore] - 移除 CheckSpaceMember 相关冗余代码

## 1.1.4
1. [fix] - fix内部 API 用户信息中间件，支持可选的账户类型检查

## 1.1.5
1. [opt] - 简化记忆检索查询参数，移除冗余的 all_inputs 拼接逻辑
2. [fix] - 修复 input_fields.go 中变量名拼写错误
3. [chore] - 模块路径迁移至 GitHub，升级 Go 版本至 1.24.0

## 1.1.7
1. [refact] - 将私有依赖路径迁移至 GitHub（proton-rds-sdk-go、proton-mq-sdk-go、TelemetrySDK-Go）
2. [chore] - 更新依赖版本：TelemetrySDK-Go v2.10.3, proton-rds-sdk-go v1.4.2, proton-mq-sdk-go v1.9.1

## 1.1.8
1. [chore] - 升级 github.com/bytedance/sonic 从 v1.14.1 到 v1.14.2
2. [chore] - 自动更新相关依赖包版本（包括 go-playground/validator、oklog/ulid、uber/mock 等）

