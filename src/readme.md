# 结构说明

## domain
- valueobject：值对象

## drivenadapter
从动适配层，实现领域模型与外部系统的交互，包括：
- dbaccess：数据库访问

## driveradapter
驱动适配层，处理外部请求并调用领域服务，包括：
- api：API接口
- mq：消息队列
- task：任务处理

## infra
基础设施层，提供技术支持，包括：
- persistence：持久化相关

## port
接口（港口），定义系统与外部交互的接口，包括：
- driven：从动端口，由领域层定义，由从动适配器实现
- driver：驱动端口，由领域层实现，由驱动适配器调用
