# Todo, やったこと管理アプリ

 自分の勉強用に作成したアプリのソースコードを公開させていただきます。  
 機能面、実装面で指摘点等あればどしどし下さい。  
 なお、勉強用なのでログインや認証機能など、自分が利用する上では不要な機能なども入っております。

## 用いた技術

- Docker

`docker-compose up` でDB(postgres11.8)含めてすぐにアプリが起動できます。ポートは5080です。（起動できなかったらお申し付けください。。)

- Go

Goの勉強がしたかったのでサーバーサイドはgoで記述しました。  
API作成は標準パッケージのnet/httpで頑張り、他gorp（ORマッパー）を利用したDB操作や、JWTの認証機能を実装しました。  
また、controller層、dao層など、レイヤーを意識して作成しました。  

- Vue

フロントサイドはVueJsを用いて開発しており、フロントとサーバーで分離した環境を意識しました.  
Vue Routerを用いてレンダリングしています。  
またサーバーとの通信はaxiosを用いており、interceptorsなどを使って共通エラー処理やヘッダー処理なども実装しました。

## 機能(`docker-compose up` で動かしていただいた方のため)

- Sign Up

アプリの利用はSign Upから行います。
(* 入力チェックなどはまだ実装していないのでぶっちゃけなんでもOKです)

- ログイン

Sign Upを終えたらログイン画面に戻るので登録したユーザーIDとパスワードを入力してログインします。  
(ログインなどをせずにhomeボタンなどを押すと認証エラーとなります。)

- Home

Home画面という名前になっていますが実際はやったこと管理機能です。
この時間は何をやったというのを記録していきます（日記の時間単位）。  
期間を範囲で指定して特定期間にやったことなどを後で見返すことも可能です。  
右上の登録アイコンボタンから記録できます。

- ToDo

Todo管理はTodoの一覧の表示と、終えたTodoに関しては下のリストに移るような仕様になっています。 
完了の判断は終了日程が埋まっているかによって判断しています。




