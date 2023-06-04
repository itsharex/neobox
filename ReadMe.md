## NeoBox
----------------
NeoBox 是一款基于sing-box和xray-core的命令行客户端。
它能自动帮你从免费vpn中筛选出可用的，并在本地启动代理。默认端口为2019。入站协议支持Sock5/Http。
出站协议支持范围和sing-box相同，主要有vmess、vless、trojan、shadowsocks、shadowsocksR等。

NeoBox提供一个交互式shell，用于操作NeoBox提供的服务。交互式shell提供命令如下：
```bash
Commands:
  add           Add proxies to neobox mannually. # 添加自己的私有vpn，如果你已有的话
  clear         clear the screen # 清屏
  exit          exit the program # 退出shell
  export        Export vpn history list. # 导出历史数据(即之前获取到的有效vpn)到json文件
  filter        Filter vpns by verifier. # 手动开启筛选(-f选项表示强制从新下载文件，-hs选项表示历史数据也加入待筛选队列之中)
  gc            Start GC manually. # 手动触发一下垃圾回收，如果闲暇时占用内存较多(待测试)。
  geoinfo       Install/Update geoip&geosite for sing-box. # 下载或更新 geoip.db以及geosite.db
  help          display help # 显示帮助信息。
  parse         Parse raw proxy URIs to human readable ones. # 将rawlist解析为人类能读懂的格式，主要用于发现proxy uri的规则。
  pingunix      Setup ping without root for Unix/Linux. # 在Linux下设置不需要root权限的ping操作。
  restart       Restart the running sing-box client with a chosen vpn. [restart vpn_index] # 从新启动一个client，使用指定的vpn_index。
  show          Show neobox info. # 显示neobox相关的信息，例如vpn统计信息，vpn可用列表，neobox运行情况，当前正在使用的vpn等等。
  start         Start an sing-box client/keeper. # 开启一个client和一个keeper。
  stop          Stop the running sing-box client/keeper. # 停止已经开启的client和keeper。
```

## 注意事项
- 如果geoinfo没有下载成功，需要手动下载，否则neobox不能成功运行；
- 如果是Linux环境，需要使用pingunix获取免root执行ping的权限；

## 效果展示
![neobox-shell](https://github.com/moqsien/neobox/blob/main/docs/neobox.png)

## 感谢
- [sing-box](https://github.com/SagerNet/sing-box)
- [xray-core](https://github.com/XTLS/Xray-core)
