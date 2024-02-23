## 演示视频

https://www.bilibili.com/video/BV1dx421Z7qu

## 项目目录

```
douyin
|-- cmd
	|-- hertz-cmd              包含hertz启动文件
	|-- kitex-cmd              包含kitex启动文件
	|-- mq-cmd                 包括消息队列服务启动文件
|-- idl
	|-- hertz-idl              生成hertz的thrift文件
	|-- kitex-idl              生成kitex的thrift文件
|-- internal
	|-- hertz-server
		|-- handler            api网关的业务代码
		|-- model              hertz的生成代码
		|-- router             hertz的生成代码
	|-- kitex-server
		|-- handler            rpc服务业务代码
	|-- pkg
		|-- dal                数据库操作
		|-- kitex_client       封装了kitex的rpc请求
		|-- kitex_gen          kitex的生成代码
		|-- mw                 中间件
		|-- mq                 消息队列
		|-- parse              解析配置文件
		|-- logger             日志文件操作
|-- deployment
	|-- config                 配置文件
|-- doc
|-- go.mod
|-- go.sum
```

## 实现主要功能

|                | **互动方向**                                                 | **社交方向**                                             |                                                              |
| -------------- | ------------------------------------------------------------ | -------------------------------------------------------- | ------------------------------------------------------------ |
| 基础功能项     | 视频 Feed 流、视频投稿、个人主页                             |                                                          |                                                              |
| 基础功能项说明 | 视频Feed流：支持所有用户刷抖音，视频按投稿时间倒序推出视频投稿：支持登录用户自己拍视频投稿个人主页：支持查看用户基本信息和投稿列表，注册用户流程简化 |                                                          |                                                              |
| 方向功能项     | 喜欢列表                                                     | 用户评论                                                 | 关系列表                                                     |
| 方向功能项说明 | 登录用户可以对视频点赞，在个人主页喜欢Tab下能够查看点赞视频列表 | 支持未登录用户查看视频下的评论列表，登录用户能够发表评论 | 登录用户可以关注其他用户，能够在个人主页查看本人的关注数和粉丝数，查看关注列表和粉丝列表 |

## 架构设计
![Pasted image 20240218181454.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/1db6b709cc7a4d40b91c28bf83c2a3e5~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=993&h=408&s=31903&e=png&b=fefefe)
## 微服务拆分
![Pasted image 20240219112924.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/23257aca4bf543909f6486b148ff0b08~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=775&h=622&s=37040&e=png&b=fdfdfd)

## 性能测试

### 测试流程
- 每个接口在短时间内请求50次，并且时间间隔为0
- 测试代码如下
```javascript
pm.test("success send", function (){
    const responseJson = pm.response.json();
    pm.expect(responseJson.status_code).to.eql(0);
})
```
### 发布评论
- 平均响应速度为83ms，50个测试用例全部通过
![Pasted image 20240223121348.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/41ee8a71f47345b59c80ebd4e7b747d2~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1422&h=887&s=169742&e=png&b=fdfdfd)
### feed流
- 平均响应速度为347ms，50个测试用例全部通过
![Pasted image 20240223121715.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/77bf5d8c63d24b6b9a9c6cee1d374d80~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1407&h=890&s=163813&e=png&b=fdfdfd)
### 查看视频评论
- 平均响应速度为470ms，50个测试用例全部通过
![Pasted image 20240223122001.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d9300f7d25cc494692d78321d129ac5d~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1431&h=897&s=109489&e=png&b=fefefe)
### 查看用户信息
- 平均响应速度为33ms，50个测试用例全部通过
![Pasted image 20240223122258.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/00da98e8dc184c41a012fd90eaef6535~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=1432&h=899&s=178223&e=png&b=fefefe)
