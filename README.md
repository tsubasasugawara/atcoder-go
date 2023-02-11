# vimの設定
```bash
$ mkdir .vimsession
```
その後、vimで:PlugInstallを実行

# atcoder-cliとonline-judge-toolsにログイン
```bash
$ acc login
? username: [username]
? password: [password]
OK

$ oj login https://atcoder.jp/
Username: [Username]
Password: [Password]
```

# テストケース取得
```bash
$ acc new [contestID] #contestIDはABCの288回の場合はabc288
```

# テスト
```bash
# テストしたい問題のフォルダに入る
$oj t
```

# コードの提出
```bash
$acc s
Are you sure? Please type ""
```
