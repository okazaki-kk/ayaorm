# ayaorm

Ayaormは自動生成を用いたActive Record風のインターフェイスを持つORMです。

## How To Install

自動生成に用いるバイナリは以下のコマンドで取得できます。

```shell
go install github.com/okazaki-kk/ayaorm/cmd/ayaorm@latest
```

また、goのコード内でayaormを使用するには、以下のコマンドでダウンロードできます。
```
go get -u github.com/okazaki-kk/ayaorm
```

## Quick Start

まず、User構造体を含んだ`schema.go`ファイルを作成します。

```go
package main

import "github.com/okazaki-kk/ayaorm"

type User struct {
 ayaorm.Schema
 Name string
 Age  int
}
```

そして、
```shell
ayaorm schema.go
```
を実行すると、`schema_gen.go`というファイルが自動生成されます。
これで準備完了です。`User`構造体の中に`ayaorm.Schema`が定義されていますが、これを含めることで、primarykey idとcreated_at, updated_atが自動的に追加されます。

後は以下のように簡潔にDB操作を行えます。(sqlite3を用いた例)

```go
package main

import (
 "database/sql"
 "fmt"
 "log"

 _ "github.com/mattn/go-sqlite3"
)

func main() {
  // insert into users (name, age) values ('example-user', 16)
  user, err := User{}.Create(UserParams{Name: "example-user", Age: 16})

  // select count(*) from users;
  count := User{}.Count()

  // select * from users where age = 16;
  user, err = User{}.Where("age", 16).Query()

  // select * from users where age > 15 or name = 'example-user';
  user, err = User{}.Where("age", ">", 15).Or("name", "example-user").Query()

  // update users set age = 35 where id = xxx;
  err = user.Update(UserParams{Age: 35})
}
```

## モデルの操作
まず、基本的なCRUD操作の方法を紹介します。
### Create
```go
// insert into users (name, age) values ('example-user', 16)
user, err := User{}.Create(UserParams{Name: "example-user", Age: 16})

// insert into users (name, age) values ('example-user', 16)
user = User{Age: 16, Name: "example-user"}
err = user.Save()
```

また、`CreateAll`メソッドを用いると、複数のレコードを一度に挿入することができます。
```go
// insert into users (name, age) values ('example-user', 16), ('example-user2', 17)
users, err := User{}.CreateAll([]UserParams{
  {Name: "example-user", Age: 16},
  {Name: "example-user2", Age: 17},
})
```

### Update
```go
// update users set age = 35 where id = xxx;
err = user.Update(UserParams{Age: 35})
```

### Delete
```go
// delete from users where id = xxx;
err = user.Delete()
```

## クエリ
次に、データベースを検索するクエリを紹介します。
WhereやOrderByなどはQuery()メソッドを呼ばないとクエリが叩かれないため、チェーンメソッドとして使用できます。一方、FirstやLastなどはクエリが即時に実行されるため、チェーンメソッドとして使用することはできません。

### Where
```go
// select * from users where age = 16;
user, err = User{}.Where("age", 16).Query()

// select * from users where age > 15 and name = 'example-user';
user, err = User{}.Where("age", ">", 15).And("name", "example-user").Query()

// select * from users where age > 15 or name = 'example-user';
user, err = User{}.Where("age", ">", 15).Or("name", "example-user").Query()

// select * from users where name is null;
user, err = User{}.Where("name", nil).Query()

// select * from users where age > 15 and (name = 'example-user' or name = 'example-user2');
user, err = User{}.Where("age", ">", 15).And(func(q *Query) {
  q.Where("name", "example-user").Or("name", "example-user2")
}).Query()

// select * from users where age in (15, 16, 17);
user, err = User{}.Where("age", []int{15, 16, 17}).Query()
```

### find
```go
// select * from users where id = xxx limit 1;
user, err = User{}.Find(xxx)

// select * from users where name = "example-user" limit 1;
user, err = User{}.FindBy("name", "example-user")

// select * from users order by id asc limit 1;
user, err = User{}.First()

// select * from users order by id desc limit 1;
user, err = User{}.Last()

// select * from users;
users, err = User{}.All()
```

### その他
```go
// select count(*) from users;
count = User{}.Count()

// select * from users order by age desc;
users, err = User.OrderBy("age", "desc")

// select * from users group by name having count(name) = 2;
User{}.GroupBy("name").Having("count(name)", 2).Query()
```

## バリデーション
schemaファイルにバリデーションを記述することで、値が不正な場合にデータベースにアクセスする前にエラーを返すことができます。

例えば、Userモデルのageカラムを以下のように定義した場合、ageには0以上の値しか入れることができません。
```go
// schema.go
func (m User) validateNumericalityOfAge() Rule {
	return MakeRule().Numericality().Positive()
}
```
このように、validateNumericalityOf***から始まるメソッドは、数値に関するバリデーションであると認識されます
また、xxxにはバリデーション対象のカラム名を、レシーバにはバリデーション対象の構造体を指定します。

この場合、ageカラムに負の値が指定されているとデータベースにアクセスせずにエラーを返します。また、現在の状態がバリデーションに引っかからないかは、`IsValid()`メソッドで判定できます。
```go
user := User{Age: -1}
err := user.Save() // insert into ~~~ は実行されない
// err = "Age must be positive"

ok, errors = user.IsValid()
// バリデーションエラーが発生した場合、okはfalse、errorsがエラー配列
```

### 数値バリデーション
```go
// ageは負の値のみ許可
func (m User) validateNumericalityOfAge() Rule {
  return MakeRule().Numericality().Negative()
}
```

### 文字列バリデーション
`validateLengthOf***`から始まるメソッドは、string型のカラム***の長さに関するバリデーションであると認識されます。
```go
// Userのnameカラムは、3文字以上10文字以下
func (m User) validateLengthOfName() Rule {
	return MakeRule().MaxLength(10).MinLength(3)
}
```

### OnCreateとOnUpdateフック
以下のようにメソッドチェーンを用いて、バリデーションを行うタイミングを指定することができます。
```go
// Userのnameカラムバリデーションは、モデルの作成時のみ有効
func (m User) validateLengthOfName() Rule {
	return MakeRule().MaxLength(10).MinLength(3).OnCreate()
}
```

```go
// Userのageカラムバリデーションは、モデルの更新時のみ有効
func (m User) validateNumericalityOfAge() Rule {
  return MakeRule().Numericality().Negative().OnUpdate()
}
```

### バリデーションのカスタマイズ
`CustomRole`メソッドを用いると、バリデーションのルールをカスタマイズすることができます。
以下の例では、nameカラムに"custom-example"という値が入っている場合にエラーを返すようにしています。
```go
func (m User) validateCustomRule() validate.Rule {
	return validate.CustomRule(func(es *[]error) {
		if m.Name == "custom-example" {
			*es = append(*es, errors.New("name must not be custom-example"))
		}
	})
}
```

## リレーション
モデル間の関連付けを行うことで、自分のコードの共通操作がシンプルになって扱いやすくなります。

例えば、以下のPostモデルとCommentモデルを考えます。
```go
type Comment struct {
	ayaorm.Schema
	Content string
	Author  string
	PostId  int
}

type Post struct {
	ayaorm.Schema
	Content string
	Author  string
}
```
そして、CommentとPostが1対多の関係にあるとします。この場合、Postモデルに以下のようにHasManyメソッドをschemaファイルに定義します。
```go
// schema.go
// メソッド名はhasMany***sでなければならない
func (m Post) hasManyComments() {}
```

すると、自動生成されたファイルによって、`Comments()`というメソッドが使えるようになり、関連するCommentモデルの配列を取得することができます。また、`JoinComments`というメソッドも使えるようになり、関連するCommentモデルを結合したクエリを取得することができます。
```go
// select * from comments where post_id = xxx;
post, _ = Post{}.Find(xxx)
comments, _ := post.Comments()

// select * from posts join comments on posts.id = comments.post_id where posts.id = xxx;
Post{}.JoinComments().Find(xxx)
```

一方、Commentモデルには以下のようにBelongsToメソッドを定義します。
```go
// schema.go
// メソッド名はbelongsTo***でなければならない
func (m Comment) belongsToPost() {}
```

すると、`Post()`というメソッドが使えるようになり、関連するPostモデルを取得することができます。また、`JoinPost`というメソッドも使えるようになり、関連するPostモデルを結合したクエリを取得することができます。
```go
// select * from posts where id = xxx;
comment, _ = Comment{}.Find(xxx)
post, _ := comment.Post()

// select * from comments join posts on comments.post_id = posts.id where comments.id = xxx;
Comment{}.JoinPost().Find(xxx)
```

他にも、hasOneのようなメソッドを定義することで、1対1の関係を表現することもできます。

## null
nullを扱う場合は、以下のように`null.Null***`型で対応できます。この場合、Addressカラムはnullを許容するようになります。
```go
type User struct {
	ayaorm.Schema
	Name    string
	Age     int
	Address null.NullString
}
```
