
---
label: til
layout: default
title: ActiveRecord fun dog (yaoming)
---
## joins
```ruby
User.joins(:posts)
#=> Select "users".* FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id"
```
You can use string in order to customize your joins
```ruby
User.joins("LEFT JOIN bookmarks ON bookmarks.bookmarkable_type = 'Post' AND bookmarks.user_id = user.id")
#=> SELECT from users LEFT JOIN bookmarks ON bookmarks.bookmarkable_type = 'Post' AND bookmarks.user_id = users.id
```
## where
### string
```ruby
Client.where("orders_count = '2'")
#=> SELECT * FROM clients where orders_count = 2
```
SQL Injections ! It's recommended to use one of following methods 
### array
```ruby
User.where ["name = ? and email = ? ", "Joe", "joe@framgia.com"]
#=> SELECT * FROM users where users.name = 'Joe' AND users.email = 'joe@framgia.com'
```
`ActiveRecord` takes care of building query to avoid SQL injection
### hash
Fields can be symbols or strings. Values can be single values, arrays, or ranges
```ruby
User.where {name: "Joe", email: "joe@framgia.com"}
#=> SELECT * FROM users WHERE name = 'Joe' AND email = 'joe@framgia.com'

User.where {name: ["Alice", "Bob"]}
#=> SELECT * FROM users WHERE name IN ('Alice', 'Bob')

User.where {age: (1...10)}
#=> SELECT * FROM users WHERE age BETWEEN 1 AND 10
```
### Joins
```ruby
User.joins(:posts).where "posts.created_at < ?", Time.now

User.joins(:posts).where {"posts.published": true}

User.joins(:posts).where {posts: {published: true}}
```
## select
Works in two unique ways
1. takes a block so it can be used just like `Array#select`
```ruby
Model.all.select {|m| m.field == value}
```
2. Modifies the SELECT statement for the query so that only certain fields are retrieved.
```ruby
Model.select :field
#=> [#<Model id: nil, field: "value>]
```
The agurment to the method can also be an array of fields
```ruby
Model.select :field, :other_field, :and_one_more
#=> [#<Model id: nil, field: "value", other_field: "other_value", and_one_more: "and one more">]
```
You can also use one or more strings, which will be used unchanged as SELECT fields.
```ruby
Model.select "field AS field_one", "other_field AS field_two"
#=> [#<Model id: nil, field: "value", other_field: "other_value">]
Mode.select("field as field_one").first.field 
#=> "value"
```
## [rewhere](https://apidock.com/rails/v4.2.7/ActiveRecord/QueryMethods/rewhere)
## preloads 
Allows preloading, in the same way that `includes` does: 
```ruby
User.preload :posts
#=> SELECT "posts".* FROM "posts" WHERE "posts"."user_id" IN (1, 3)
```
## Order
Allows to specify an `order` attribute
```ruby
User.order :name
#=> SELECT "users".* FROM "users" ORDER BY "users"."name" ASC

User.order email: :desc
#=> SELECT "users".* FROM "users" ORDER BY "users"."name" DESC

User.order :name, email: desc
#=> SELECT "users".* FROM "users" ORDER BY "users"."name" ASC, "users"."email" DESC
```
## offset
```ruby
User.offset 10
User.order(:name).offset 10 
```
## limit 
Specifies a limit for the number of record returned
```ruby
User.limit 10
#=> ...LIMIT 10
```
## includes
Specifies relationships to be included in the result set
```ruby
users = User.includes :address 
users.each do |user|
  user.address.city
end 
```
allows you access the address attributes without firing any addition query. This will often result in a performance improvement over a simple join.
You can include multiple relationships:
```ruby
User.includes :orders, :address
```
### condition 
```ruby
User.includes(:posts).where "posts.name = ?", "example"
```
Note: `includes` works with associations name while `references` works with actual table name

## having
Allows to specify a HAVING clause. Can't work without `group`
```ruby
Order.having("SUM(price) > 30").group :user_id
```
## group
Allows to specify `group` attribute
```ruby
User.group :name
#=> SELECT "users".* FROM "users" GROUP BY "users"."name"
```
Returns an array with distinct records based on the group attribute
```ruby
User.select :id, :name
#=> [#<User id: 1, name: "Oscar">, #<User id: 2, name: "Oscar">, #<User id: 3, name: "Foo">]
User.group :name

```








