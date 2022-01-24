### Kratos

##### 项目规范
1,每个目录 需要有独立的README.md  CHANGELOG.md CONTRIBUTORS.md，具体可以参考：
http://git.bilibili.co/platform/go-common/tree/master/business/serv


ice/archive

2,以后每个业务或者基础组件维护自己的版本号，在CHANGELOG.md中，rider 构建以后的tag关联成自己的版本号；

3,整个大仓库不再有tag，只有master 主干分支，所有mr发送前，一定要注意先merge master；

4,使用Rider构建以后（retag），回滚可以基于Rider的retag来回滚，而不是回滚大仓库的代码；

5,提供RPC内部服务放置在business/service中，任务队列放置在business/job中，对外网关服务放置在business/interface，管理后台服务放置在business/admin

6,每个业务自建cmd文件夹,将main.go文件和test配置文件迁移进去

7,构建的时候自定义脚本选择krotos_buil.sh,自定义参数选择自己所在业务的路径 （ps：例如 interface/web-show）

8,大仓库的mr合并方式为，在mr中留言"+merge"，鉴权依据服务根目录下 CONTRIBUTORS.md 文件解析，具体可以参考：
http://info.bilibili.co/pages/viewpage.action?pageId=7539410