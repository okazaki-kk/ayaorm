# ayaorm

Ayaorm is an ORM with an Active Record-like interface using auto-generation.

## How To Install

Binaries for automatic generation can be obtained with the following command.

```shell
go install github.com/okazaki-kk/ayaorm/cmd/ayaorm@latest
```

Also, to use ayaorm in your go code, you can download it with the following command

```
go get -u github.com/okazaki-kk/ayaorm
```

## Quick Start

First, create a `schema.go` file containing the User structure.

```go
package main

import "github.com/okazaki-kk/ayaorm"

type User struct {
 ayaorm.Schema
 Name string
 Age  int
}
```

Then
```shell
ayaorm schema.go
```

a file named `schema_gen.go` will be automatically generated.
This completes the preparation. Schema` is defined in the `User` structure, and by including it, the primarykey id and created_at, updated_at are automatically added.

The rest of the DB operations can be performed in a concise manner as follows. (Example using sqlite3)

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

## Model Manipulation
First, here are some basic CRUD operations.
### Create
```go
// insert into users (name, age) values ('example-user', 16)
user, err := User{}.Create(UserParams{Name: "example-user", Age: 16})

// insert into users (name, age) values ('example-user', 16)
user = User{Age: 16, Name: "example-user"}
err = user.Save()
```

The `CreateAll` method can also be used to insert multiple records at once.
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

## Query
Next, we introduce queries that search the database.
Where, OrderBy, etc. can be used as chained methods because the query is not struck until the Query() method is called. On the other hand, First, Last, etc. cannot be used as chained methods because the query is executed immediately.

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

### Others
```go
// select count(*) from users;
count = User{}.Count()

// select * from users order by age desc;
users, err = User.OrderBy("age", "desc")

// select * from users group by name having count(name) = 2;
User{}.GroupBy("name").Having("count(name)", 2).Query()
```

## Validation
By writing validation in the schema file, an error can be returned before accessing the database if the value is incorrect.

For example, if the age column of the User model is defined as follows, age can only contain values greater than or equal to 0.
```go
// schema.go
func (m User) validateNumericalityOfAge() Rule {
	return MakeRule().Numericality().Positive()
}
```
Thus, methods beginning with validateNumericalityOf*** are recognized as validation on numeric values
The xxx is the name of the column to be validated, and the receiver is the structure to be validated.

In this case, if a negative value is specified for the age column, an error is returned without accessing the database. You can also use the `IsValid()` method to determine if the current state is not trapped by validation.
```go
user := User{Age: -1}
err := user.Save() // insert into ~~~~ is not executed
// err = "Age must be positive"

ok, errors = user.IsValid()
// If a validation error occurs, ok is false and errors is the error array
```

### Numerical Validation
```go
// Only negative values are allowed for age
func (m User) validateNumericalityOfAge() Rule {
  return MakeRule().Numericality().Negative()
}
```

### String Validation
Methods beginning with `validateLengthOf***` are recognized as validations on the length of a column*** of type string.
```go
// User's name column must be between 3 and 10 characters long
func (m User) validateLengthOfName() Rule {
	return MakeRule().MaxLength(10).MinLength(3)
}
```

### OnCreate and OnUpdate hook
The following method chain can be used to specify when to perform validation.
```go
// User's name column validation is valid only when creating the model
func (m User) validateLengthOfName() Rule {
	return MakeRule().MaxLength(10).MinLength(3).OnCreate()
}
```

```go
// User's name column validation is valid only when updating the model
func (m User) validateNumericalityOfAge() Rule {
  return MakeRule().Numericality().Negative().OnUpdate()
}
```

### Customize Validation
The `CustomRole` method allows you to customize the validation rules.
In the following example, an error is returned if the name column contains the value "custom-example".
```go
func (m User) validateCustomRule() validate.Rule {
	return validate.CustomRule(func(es *[]error) {
		if m.Name == "custom-example" {
			*es = append(*es, errors.New("name must not be custom-example"))
		}
	})
}
```

## Relation
By making associations between models, the common operations of your code become simpler and easier to handle.

For example, consider the following Post and Comment models.
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
And suppose that Comment and Post have a one-to-many relationship. In this case, define the HasMany method in the Post model in the schema file as follows.
```go
// schema.go
// Method name must be hasMany***s
func (m Post) hasManyComments() {}
```

Then, the auto-generated file allows you to use the method `Comments()` to retrieve an array of related Comment models. Also, the method `JoinComments` can be used to retrieve a query that joins related Comment models.
```go
// select * from comments where post_id = xxx;
post, _ = Post{}.Find(xxx)
comments, _ := post.Comments()

// select * from posts join comments on posts.id = comments.post_id where posts.id = xxx;
Post{}.JoinComments().Find(xxx)
```

On the other hand, the Comment model defines the BelongsTo method as follows.
```go
// schema.go
// Method name must be belongsTo***
func (m Comment) belongsToPost() {}
```

Then the method `Post()` can be used to retrieve the related Post models. Also, the method `JoinPost` can be used to retrieve a query that joins related Post models.
```go
// select * from posts where id = xxx;
comment, _ = Comment{}.Find(xxx)
post, _ := comment.Post()

// select * from comments join posts on comments.post_id = posts.id where comments.id = xxx;
Comment{}.JoinPost().Find(xxx)
```

Other one-to-one relationships can also be expressed by defining methods such as hasOne.

## null
When dealing with null, the `null.Null***` type can be used as follows. In this case, the Address column will allow null.
```go
type User struct {
	ayaorm.Schema
	Name    string
	Age     int
	Address null.NullString
}
```
