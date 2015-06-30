Go言語でWebアプリ開発する雛形です.


# 大事なファイル

## ```watch.sh```

これを動作させるとJSXファイル/lessファイル/htmlファイルの変更を検出すると、

1. jsxコンパイル
2. lessコンパイル
3. js圧縮
4. css圧縮
5. htmlコピー

を自動で行うようになります.

## bindata.sh

静的リソース(要はGoのソースファイル以外)をGoのソースコードに埋め込みます.

ソースは以下のファイルに出力されます.

```
src/bindata/bindata.go
```

```
src/bindata/debug/bindata_debug.go
```


debugモードでは静的リソースに対するリクエストの度に、ディスクからファイルを読み込みます.  
debugモードの切替は環境変数で行います(```production/debug```).

静的リソースを追加した場合、```bindata.sh```を実行し、Webアプリを再起動する必要があります.

# Webアプリの起動

```
go run src/main.go -bind=:<port>
```

