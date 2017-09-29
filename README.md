### なんのツールなのか
catworkはChatworkへの通知用コマンドです．標準入力で値を受け取って
Chatworkに通知します．
# 環境構築
### GOPATHの指定
```
export GOPATH=$HOME/.go
```
### API key とROOMIDの設定ファイルへの設定の外出しは後からやります(bow) 
### 使い方
```
cat * |go run catwork.go -r "*ROOM ID*" -k "IP Key"
```
