### 名称： ITflow

# 有点忙，  不忙的时候， 会继续更新 内容
- [ ] go代码优化
- [ ] vue代码优化
- [ ] 增加docker-compose一条命令启动服务
- [ ] 支持分布式部署

### 简介
  一个开源的bug管理系统，IT人员开发全过程，文件存储，接口文档, 单接口测试功能

文档地址： https://www.hyahm.com/article/257  

### 展示页面： 
   展示页面会更新为最新可使用的代码  
   [ITflow](http://bug.hyahm.com "ITflow")  
   
 

### 项目优势   
 1， 部署简单,使用简单    
 2， 因为后端是go语言写的，跨平台，速度快（虽然这个没什么卵用的样子）  
 3， vue平台用的vue-element-admin框架上写的  
 4， 有需求就有更新  
 5， 永久开源  可以自己二次开发  
 6， 会有专门的接口文档，方便适应各种后端语言  
 7， 拥有api接口文档  
 8， 增加共享文件夹  
 9， 数据库表自动更新，不影响使用的情况随时升级  
 
###   功能完成  
  1， 增加bug，改变bug状态，转交bug  
  2， 显示bug列表,搜索、分页  
  3， 用户创建及其操作  
  4， 上传个人头像  
  5， 添加部门  
  6， 增加缓存机制   
  7， 增加邮件通知功能  
  8， 增加admin用户的信息重置接口  
   admin用户有且只有一个，注册admin账户建议直接操作数据库，然后修改密码即可
   如果忘记admin的密码，可以执行下面命令重置密码，如下所示，只能在go服务器那台机器上执行
```
   curl http://127.0.0.1:10001/admin/reset?password=123
```
  9，  增加修改邮箱，昵称，姓名页面  
  10， 只允许修改自己部门的账号权限   增加用户修改权限功能  
  11， bug可以指定多人，自己的bug才可以转交，删除bug内部转交功能，增加缓存,增加查看所有bug的权限  
  12， 增加用户禁用功能，当此用户存在bug时，无法被删除  
  13， 禁用用户，此用户的所有发布的bug都将移动至垃圾箱，垃圾箱里面的bug只有管理员才能查看，启用用户会将此用户的bug改为非垃圾箱  
  14， 增加操作日志，只有管理员才能查看   
  15， 状态实时保存，增加缓存  
  16， 数据库表自动更新
  17,  接口文档自定义数据类型， 无法定义基础类型  
  18， 增加接口测试  
  
### QQ群  
    928790087@qq.com  
