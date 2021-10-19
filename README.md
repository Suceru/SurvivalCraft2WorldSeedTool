# SurvivalCraft2WorldSeedTool
You can convert the original seed into a unique seed number for the map, or you can multiply the seed number to get numerous seeds.
## 使用
The procedure for obtaining a seed number is as follows:

```
1. Unpack the.scworld file or look in the "LocalState\Worlds\" folder for the relevant folder.
2. Locate <Value Name="WorldSeed" Type="int" Value="49" /> in Project.xml.
3. where the seed number is 49 and Value="49" is the value.
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
## Attribution
- Code
    - Suceru
## 额外
当git拉取失败时，可以试一下这段命令。 
```
git config --global http.sslVerify "false"
git config --global --unset http.proxy
git config --global --unset https.proxy
```
## License
The GPL is used to license this project.
