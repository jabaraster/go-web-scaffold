Go言語でWebアプリ開発する雛形です.


# 事前準備

下記ソフトウェアをインストールしておくこと.

* Go言語: サーバサイドはGo言語で実装
* node.js: クライアント開発環境
* npm: node.jsのライブラリ管理

# フォルダ構成

___ 編集するファイルは以下のファイルです ___

* ```src```ディレクトリ以下
  * ```src/go```：サーバサイドを実装するGoのソースディレクトリのルート
  * ```src/html```：HTML.
    編集したらgulpにより```assets```直下にコピーされる(フォルダ構造は保たれる).
    ※コピー先は変えるかも.
  * ```src/jsx```：JSX(JavaScriptのreact.js向け拡張)を格納.  
    編集したらgulpによりJSXコンパイル→require解決→圧縮→が施された上で```assets/js```にコピーされる(フォルダ構造は保たれる)
  * ```src/less```：Less(CSSを効率良く書ける言語)を格納.  
    編集したらgulpによりlessコンパイル→圧縮が施された上で```assets/css```にコピーされる(フォルダ構造は保たれる).  
    なお、include用に、先頭が ```_``` で始まるファイルはgulpから無視される.
* ```gulpfile.js```：node.js製のビルドツール__gulp__の設定ファイル.
* ```package.json```：node.jsのライブラリを記述. 手動で編集することはほとんどないはず.
* その他：フォントなど、サードパーティ製の物は```assets```以下に適切なディレクトリを掘って置く.  
  代表例はbootstrapのCSSやフォントファイル.  
  ※サードパーティ製のJSは、jsxファイルの中にrequiredして欲しいので、```assets/js```の下に置くことは極力しない.


# クライアントサイドの開発

## 必要ライブラリのインストール

まずはnpmで必要モジュールをインストール.

```
npm install
```

このコマンドで、```package.json``` の中に書かれているモジュールが全てインストールされます.

## ビルドツールの起動

```
gulp
```

gulpはnode.js製のビルドツール.  
上記コマンドを打つと```gulpfile.js```の設定に従って、ファイルを監視→ビルドを実行してくれます.  

# サーバサイドの開発

Go言語で開発します.  

## ディレクトリ構成

* ```src/go/```：```main.go```、```main_test.go```のみ格納.
* ```src/env/```：envパッケージを格納. 環境変数周りの処理を実装.
* ```src/model/```：modelパッケージを格納. DBテーブルに関連付くstructの定義、DBアクセス処理を実装.
* ```src/web/```：webパッケージを格納. Web特有の処理を実装. あまり出番はないかな. 
* ```src/web/handler/```：web/handlerパッケージを格納. クライアントからのリクエストを受け付けるメソッド(=handler)を実装.

その他、拡張は適宜検討.



# Webアプリの起動

```
go run src/go/main.go -bind=:<port>
```

jsx,less,htmlの変更はブラウザのリロードで反映されます.  
Goの変更はWebアプリを再起動しないと反映されません.  

# 技術情報

## React.js関連

__React.js__はクライアントJavaScriptのフレームワーク.

* シンプルで覚えやすい上にコードがすっきりする
* 画面のコンポーネント化が進む
* Facebookが開発／採用しており実績充分

----
* [本家](http://facebook.github.io/react/)  
  [チュートリアル](http://facebook.github.io/react/docs/tutorial.html)は一度さらっておいた方が良い.
* [Qiitaの良記事](http://qiita.com/advent-calendar/2014/reactjs)  
  よくまとまってる.

## Less関連

__Less__はCSSを効率良く記述するための言語.  

* 入れ子/変数/関数定義などの機能により冗長性を減らせる
* 便利な組み込み関数群(例えば『ちょっと暗い色』みたいな関数がある)

CSSは極端に冗長なので、Less(のような動的スタイルシート言語)を使わない開発は__悪__だとさえ思います.  

----
* [本家](http://less-ja.studiomohawk.com)  
  このページだけで、重要な機能の説明は網羅されている.  

## gorm関連

__gorm__はGo言語のDBアクセスライブラリ.  

* 高機能かつ多機能っぽい
* エンティティの関連が扱える
* 単純なクエリを素早く書ける

ただし、ときどきドキュメントにウソがあるので、そこは要注意.


----

* [GoDoc(APIドキュメント)](https://godoc.org/github.com/jinzhu/gorm)
* [GitHubリポジトリ](https://github.com/jinzhu/gorm)
* [参考）Go言語でのDBアクセスライブラリ比較](http://qiita.com/umisama/items/c022b16101c48ffdbc6a)