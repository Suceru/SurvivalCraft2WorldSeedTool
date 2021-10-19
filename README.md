# 生存战争种子工具
将原本种子变为地图唯一的种子号，也可以通过种子号计算出多个种子。
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
```
## 许可
此项目使用了GPL作为许可。
