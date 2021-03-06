# 生存战争种子工具
## 语言
:globe_with_meridians:
**简体中文**,
[English][EN],
[Bahasa Indonesia][ID],
[Català][CA],
[Deutsch][DE],
[Español][ES],
[Français][FR],
[Italiano][IT],
[Português][PT_BR],
[Русский][RU],
[العربية][AR],
[हिंदी][HI_IN],
[正體中文][ZH_TW],
[한국어][KO_KR]

[AR]:README.ar.md
[ID]:README.id.md
[CA]:README.ca.md
[DE]:README.de.md
[EN]:README.md
[ES]:README.es.md
[IT]:README.it.md
[FR]:README.fr.md
[PT_BR]:README.pt_br.md
[ZH_TW]:README.zh_tw.md
[ZH_CN]:README.zh_CN.md
[KO_KR]:README.ko_kr.md
[HI_IN]:README.hi_in.md
[RU]:README.ru.md

将原本种子变为地图唯一的种子号，也可以通过种子号计算出多个种子。
## Download
[![Downloads](https://img.shields.io/badge/2.21%20MB-Download-brightgreen)](https://github.com/Suceru/SurvivalCraft2WorldSeedTool/releases/latest)
## 使用
种子号获取方式

```
1. 从.scworld文件解压缩包或者找到“LocalState\Worlds\”中相应的文件夹。
2. 打开Project.xml，找到<Value Name="WorldSeed" Type="int" Value="49" />。
3. 其中Value="49"就是种子号。
```

打开种子工具之后，输入数字按回车选择相应功能。

```
功能
1. 种子号转为种子
2. 种子转为种子号
3. 通过种子号显示多种种子的组合方式
4. 帮助
```
## 项目
此项目使用了Golang作为语言进行编译（推荐使用Goland进行管理）。
## 贡献
- 代码
    - Suceru
## 额外
当git拉取失败时，可以试一下这段命令。 
```
git config --global http.sslVerify "false"
git config --global --unset http.proxy
git config --global --unset https.proxy
```
## 许可
此项目使用了GPL作为许可。
