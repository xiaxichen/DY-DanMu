# DY-DanMu
# 斗鱼弹幕抓取分析程序
## 具体架构
![架构.png](https://github.com/xiaxichen/DY-DanMu/blob/master/doc/%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%9E%B6%E6%9E%84%E5%9B%BE.png)

## 具体思路
根据斗鱼开发文档 https://open.douyu.com/source/api/63
进行websocket端抓取 
![协议.png](https://github.com/xiaxichen/DY-DanMu/blob/master/doc/%E5%8D%8F%E8%AE%AE.png)

消息长度：4 字节小端整数，表示整条消息（包括自身）长度（字节数）。
消息长度出现两遍，二者相同。
消息类型：2 字节小端整数，表示消息类型。取值如下：
689 客户端发送给弹幕服务器的文本格式数据
690 弹幕服务器发送给客户端的文本格式数据。
加密字段：暂时未用，默认为 0。保留字段：暂时未用，默认为 0。
websocket标准请查询
https://datatracker.ietf.org/doc/rfc6455/
中文版讲解
https://segmentfault.com/a/1190000005680323

###具体服务为三个 
    1. spider端:
        连接websocket后实时获取弹幕信息通过gorutine连接rpc服务器进行存储
    2. RPC-server端:
        连接数据库进行数据存储 查询工作
    3. Web端:
        实现api接口进行参数解析 连接rpc进行查询

![协议.png](https://github.com/xiaxichen/DY-DanMu/blob/master/doc/%E6%B5%81%E7%A8%8B%E5%9B%BE.png)