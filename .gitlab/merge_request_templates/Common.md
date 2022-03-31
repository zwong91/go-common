<!-- Title规范 : <基础库|服务名|doc|其他> : <标题> 
栗子：
    1. library/log : xxxxxxxx
    2. account-service : xxxxxxxx
    3. doc : xxxxxxxx
    4. 搞个大新闻 : xxxxxxxx
-->
### Version
* v1.0.0 : library/log
* v2.1.0 : account-service

### Related issue
fix #<issue number>

### Review tips
<!-- 写给 reviewer 的 tips -->
* <tip>
* <tip>

### .gitlab-ci.yml 提交限制
配置CI流程会自动使用govet进行代码静态检查、gofmt进行代码格式化检查、golint进行代码规范检查、gotest进行单元测试