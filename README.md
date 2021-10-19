# SurvivalCraft2WorldSeedTool
## Language
:globe_with_meridians:
[简体中文][ZH_CN],
**English**,
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
[한국어][KO_KR],
[日本][JA]

[AR]:README.ar.md
[ID]:README.id.md
[CA]:README.ca.md
[DE]:README.de.md
[EN]:README.md
[ES]:README.es.md
[IT]:README.it.md
[FR]:README.fr.md
[JA]:README.ja.md
[PT_BR]:README.pt_br.md
[ZH_TW]:README.zh_tw.md
[ZH_CN]:README.zh_CN.md
[KO_KR]:README.ko_kr.md
[HI_IN]:README.hi_in.md
[RU]:README.ru.md

You can convert the original seed into a unique seed number for the map, or you can multiply the seed number to get numerous seeds.
## Usage
The procedure for obtaining a seed number is as follows:

```
1. Unpack the.scworld file or look in the "LocalState\Worlds\" folder for the relevant folder.
2. Locate <Value Name="WorldSeed" Type="int" Value="49" /> in Project.xml.
3. where the seed number is 49 and Value="49" is the value.
```

Enter the number and press Enter to pick the correct function after starting "WorldSeedTool.exe."

```
Function 
1: Seed number are converted into Seeds.
2. Seeds are converted into Seed number.
3. Demonstrate how to mix several Seed based on their Seed number.
4. Help.
```
## Project
The compilation language for this project is Golang (Goland is recommended for management).
## Attribution
- Code
    - Suceru
## Additional
Try this command if a git pull fails.
```
git config --global http.sslVerify "false"
git config --global --unset http.proxy
git config --global --unset https.proxy
```
## License
The GPL is used to license this project.
