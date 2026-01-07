# agent-go-common-pkg

## 项目简介

    agent-go-common-pkg是data agent的go语言公共库。

## 名词解释

| 术语名称 | 定义 |
|---------|------|
| Agent（智能体） | 一种借助大模型、知识库、记忆、工具、数据流等多种基础功能，实现智能化的交互、决策、任务执行的自主实体，通过 Agent 配置进行创建，发布后可以被用户使用 |
| Agent App Market（智能体应用市场） | 一个用于查找、使用 Agent App 的入口，支持按照空间、分类等维度显示 Agent App，也可以查看历史使用的 Agent App |
| Agent Template（智能体模板） | 基于 Agent 配置保存的配置模板，模板支持导入导出，可以基于模板快速创建 Agent App |
| Agent App（智能体应用） | 基于一个或多个Agent 构建的完整的应用或系统，通常结合业务逻辑、多 Agent 协作提供特定产品能力的、解决实际场景的复杂问题 |
| Agent配置 - 基本信息 | 包括 Agent 头像、名称、简介信息 |
| Agent配置 - 所属产品 | 产品用于控制 Agent 数据源范围、是否支持临时区等设置 |
| Agent配置 - 人设&指令 | 描述AI助手的角色定位、专业领域、回复风格以及技能使用规则 |
| Agent配置 - 开场白 | Agent对话开始时的首条自动回复 |
| Agent配置 - 预设问题 | 提前设置的一系列问题，用于引导用户对话 |
| Agent配置 - 输入配置 | 配置Agent的输入相关设置 |
| Agent配置 - 知识库 | 支持AnyShare文档、Wiki、知识图谱、本体作为知识来源，Agent 可以通过知识库获取各种领域知识，实现准确推理、保证对话质量、提升效率和性能 |
| Agent配置 - 临时区 | 支持在对话过程中临时保存用户输入的临时数据、上下文、中间结果等信息，以便在生成回复时参考，临时区数据具有较短的生命周期，随着会话、任务结束而清理 |
| Agent配置 - 技能 | 支持将其他 Agent 或工具作为 Agent 的技能用于完成特定任务或实现特定功能 |

## 项目结构

```
./
├── src                           # 代码目录
│   ├── drivenadapter             # 被驱动适配器
│   │   ├── cmpt                  # 组件
│   │   ├── dbaccess              # 数据库访问
│   │   ├── httpaccess            # HTTP访问
│   ├── infra                     # 基础设施层
│   │   ├── apierr                # API错误
│   │   ├── cmp                   # 组件
│   │   ├── common                # 通用工具
│   ├── port                      # 端口层 - DDD端口
│   │   ├── driven                # 被驱动端口
│   │   └── driver                # 驱动端口
```

## ddd相关概念说明（主要是ai生成）
| 英文 |缩写| 中文 |说明 |
|----|----|----|----|
|persistence object| po |持久化对象|与数据库表结构一一对应的对象，用于数据库表的增删改查操作，属于持久化层 |
|entity| eo(entity object) |实体|具有唯一标识的领域对象，即使属性相同，不同实体也被视为不同对象|
|value object| vo |值对象|没有唯一标识的领域对象，通过其属性值来识别，相同属性值的值对象被视为同一对象|
|aggregate| - |聚合|由根实体、值对象和其他实体组成的一组相关对象的集合，作为一个整体被外界访问|
|aggregate root| - |聚合根|聚合的根实体，是外部访问聚合内部对象的唯一入口|
|repository| repo |仓储|提供对聚合的持久化和查询能力，隐藏数据访问细节|
|domain service| - |领域服务|表示领域中的一个操作，这个操作不属于任何实体或值对象|
|application service| - |应用服务|协调多个领域对象完成用户用例，是应用层与领域层的边界|


## 缩写约定（ai参与生成）
| 缩写 | 全称              |说明|
|----|-----------------|----|
| da | data agent      |数据智能体|
| um | user_management |用户管理服务相关|
| cmp | component |组件|
|dto| data transfer object |数据传输对象|在不同层或组件间传输数据的简单对象，只包含数据属性和简单的获取/设置方法，不包含业务逻辑|
| rdto | request and response data transfer object |请求和响应数据传输对象|
| d2e| dto to entity |数据传输对象到实体（当dto为rdto的req时，指将rdto转换为entity）|
|ds|datasource|数据源|data agent的数据源|
|assoc|association|关联|data agent和数据集的关联|
|obj|object|对象|数据源中的对象|
|obj_id|object id|对象id|“文档”对象的唯一标识|



