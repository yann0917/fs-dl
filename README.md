# fd-dl

> 帆书(樊登读书)音频、文稿下载, web 版接口

## 特别声明

仅供个人学习使用，请尊重版权，内容版权均为帆书APP(樊登读书)所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为帆书APP(樊登读书)所有，请勿传播内容！！！

仅供个人学习使用，请尊重版权，内容版权均为帆书APP(樊登读书)所有，请勿传播内容！！！

## 特性

* [ ] 账号密码登录
* [x] 查看分类
* [x] 查看樊登讲书、非凡精读馆、李蕾讲经典列表
* [x] 查看课程列表
* [x] 查看课程下的节目列表
* [x] 樊登讲书音频下载，文稿下载
* [x] 非凡精读馆音频下载，文稿下载
* [x] 李蕾讲经典音频下载，文稿下载
* [x] 课程音频下载，文稿下载

## 安装

### 安装依赖

`fs-dl` 支持markdown文本下载，pdf下载，以及音频视频下载，请按照自己的下载需求，安装下列依赖：

#### pdf下载

* wkhtmltopdf

> 听书文稿生成 PDF 需要借助[wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

#### 视频下载

* ffmpeg

> m3u8格式的音视频需要借助 [ffmpeg](https://ffmpeg.org/) 合成

#### markdown文本下载

不需要额外安装依赖

### 使用二进制文件安装

进入[下载列表](https://github.com/yann0917/fs-dl/releases),下载对应的系统版本，下载后即可使用。

### 使用 `go` 安装

安装go，版本需大于1.18，并设置GOPATH环境变量, 并在PATH中添加$GOPATH/bin

使用如下命令安装：

`go install github.com/yann0917/fs-dl@latest`

## 使用方法

执行 `cp config.yaml.example config.yaml`

登录[网页版](https://www.dushu365.com/)，登录后，按 F12 打开控制台，点击应用(Application)，找到`localStorage`，找到`__TOKEN` 对应的值，粘贴到配置文件`config.yaml` token; 因为接口是加密的，需要填写 aesKey 才可以加解密，aesKey 自行寻找解决方法。

`./fs-dl class` 获取分类

```text
+---------+------------+---------+----------+-------+
| 业务 ID |  业务名称  | 分类 ID | 分类名称 |  IDS  |
+---------+------------+---------+----------+-------+
|       1 | 樊登讲书   |       0 | 全部     |       |
+         +            +---------+----------+-------+
|         |            |       1 | 心灵     |     5 |
+         +            +---------+----------+-------+
|         |            |      24 | 个人成长 | 10059 |
+         +            +---------+----------+-------+
|         |            |       4 | 亲子家庭 | 10054 |
+         +            +---------+----------+-------+
|         |            |       5 | 人文历史 | 10057 |
+         +            +---------+----------+-------+
|         |            |       2 | 商业财经 | 10055 |
+         +            +---------+----------+-------+
|         |            |      27 | 社科新知 | 10056 |
+         +            +---------+----------+-------+
|         |            |      29 | 健康生活 | 10058 |
+         +            +---------+----------+-------+
|         |            |       8 | 作者光临 |    11 |
+---------+------------+---------+----------+-------+
|       2 | 非凡精读   |       0 | 全部     |       |
+         +            +---------+----------+-------+
|         |            |       1 | 心灵     |     5 |
+         +            +---------+----------+-------+
|         |            |      24 | 个人成长 | 10059 |
+         +            +---------+----------+-------+
|         |            |       4 | 亲子家庭 | 10054 |
+         +            +---------+----------+-------+
|         |            |       5 | 人文历史 | 10057 |
+         +            +---------+----------+-------+
|         |            |       2 | 商业财经 | 10055 |
+         +            +---------+----------+-------+
|         |            |      27 | 社科新知 | 10056 |
+         +            +---------+----------+-------+
|         |            |      29 | 健康生活 | 10058 |
+         +            +---------+----------+-------+
|         |            |     142 | 问书     | 10112 |
+         +            +---------+----------+-------+
|         |            |     141 | 问答专区 | 10113 |
+---------+------------+---------+----------+-------+
|       3 | 李蕾讲经典 |       0 | 全部     |       |
+         +            +---------+----------+-------+
|         |            |       1 | 心灵     |     5 |
+         +            +---------+----------+-------+
|         |            |      24 | 个人成长 | 10059 |
+         +            +---------+----------+-------+
|         |            |       4 | 亲子家庭 | 10054 |
+         +            +---------+----------+-------+
|         |            |       5 | 人文历史 | 10057 |
+         +            +---------+----------+-------+
|         |            |      29 | 健康生活 | 10058 |
+---------+------------+---------+----------+-------+
```

`./fs-dl content` 查看业务分类下的内容列表, 可查看业务为【1-樊登讲书, 2-非凡精读, 3-李蕾讲经典】下的列表

```text
+---+-----------+------------+--------------------------------------------+------------+---------+------------+
| # |  课程ID   |    标题    |                    简介                    |   主讲人   | 播放量  |  上线日期  |
+---+-----------+------------+--------------------------------------------+------------+---------+------------+
| 0 | 400113891 | 抱怨的艺术 | 让不满，变圆满                             | 樊登       |  306228 | 2024-09-21 |
| 1 | 400119107 | 五代九章   | 有趣、有料、有温度，中国人不该忽略这段历史 | 樊登       | 2040414 | 2024-09-14 |
| 2 | 400113890 | 少年发声   | 你真的知道，你的孩子在想什么吗？           | 樊登       | 3645013 | 2024-09-07 |
| 3 | 400118808 | 生活的陷阱 | 如何应对人生中的至暗时刻                   | 樊登       | 4099029 | 2024-08-31 |
| 4 | 400107837 | 达摩流浪者 | 始于背包漫游，抵达心灵自由                 | 樊登       | 2996920 | 2024-08-24 |
| 5 | 400113889 | 中年觉醒   | 发现人生下半场的幸福                       | 樊登       | 5132950 | 2024-08-17 |
| 6 | 400117830 | 一生一事   | 四十年编辑生涯苦与乐，一个人就是一部出版史 | 樊登、李昕 | 3811629 | 2024-08-10 |
| 7 | 400117820 | 一如既往   | 看透慌慌张张的世界，过上从从容容的人生     | 樊登       | 5177468 | 2024-08-03 |
| 8 | 400114785 | 永远的女儿 | 是谁无声“杀死”孩子？                       | 樊登       | 3231000 | 2024-07-27 |
| 9 | 400114201 | 康熙的红票 | 透过康熙的神秘谕令，揭开中西交流往事       | 樊登       | 3228878 | 2024-07-20 |
+---+-----------+------------+--------------------------------------------+------------+---------+------------+
```

`./fs-dl dl` 下载【樊登讲书, 非凡精读, 李蕾讲经典】的音频或者文稿， 可以下载指定ID的内容，或者批量下载

```bash
Usage:
  fs-dl dl [flags]

Examples:
fs-dl dl 123 -b1 -t1

fs-dl dl -b1 -t1 -p1 -l20

Flags:
  -r, --bookReadStatus int   阅读状态: -1 全部,0 未读, 1 已读 (default -1)
  -b, --businessType int     业务: 1-樊登讲书, 2-非凡精读, 3-李蕾讲经典 (default 1)
  -t, --downloadType int     下载格式, 1-mp3, 2-视频,  3-markdown文档, 4-PDF文档, 5-思维导图jpeg (default 1)
  -h, --help                 help for dl
  -p, --pageNo int           页码 (default 1)
  -l, --pageSize int         每页数量 (default 10)
  -s, --sort int             排序: 1-最新, 2-最热 (default 1)
```

## License

[MIT](./LICENSE) © yann0917

---
