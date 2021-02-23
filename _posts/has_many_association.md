
---
label: til
layout: default
title: has_many association
---
## Methods
```ruby
class Author < ApplicationRecord
  has_many :books
end 
```
- `author.books`
- `author.books.delete`
- `author.books.destroy`
- `author.books << book`
- `author.books.delete book`
remove one or more books from the author by settings their `foreign_key` to `NULL`
```ruby
irb(main):010:0> author.books.delete book
   (0.1ms)  begin transaction
  SQL (1.3ms)  UPDATE "books" SET "author_id" = NULL WHERE "books"."author_id" = ? AND "books"."id" = 1  [["author_id", 1]]
   (5.7ms)  commit transaction
=> [#<Book id: 1, author_id: 1, name: "Your life", created_at: "2017-05-11 10:50:36", updated_at: "2017-05-11 10:50:36">]
```
In case
```ruby
class Author < ApplicationRecord
  has_many :books, dependent: :destroy
end 
```
```ruby
irb(main):005:0> a.books.delete b
   (0.1ms)  begin transaction
  SQL (0.4ms)  DELETE FROM "books" WHERE "books"."id" = ?  [["id", 2]]
   (6.3ms)  commit transaction
=> [#<Book id: 2, author_id: 1, name: "Your name 2", created_at: "2017-05-11 11:08:50", updated_at: "2017-05-11 11:08:50">]
```
- `author.books.destroy book`
remove one or more books form author.books by running destroy on each other. 
- `author.book_ids`
return array of ids of objects in collection
- `author.books.clear`
removes all objects from collection according to the strategy specified by the dependent option. If no option is defined, it allows the default strategy. The default strategy for has_many :through is `delete_all`, 
and for `has_many` associations is set the `foreign_key` to `NULL`
```ruby
class Author < ApplicationRecord
  has_many :books 
end
```
```ruby
 a.books.clear
  SQL (0.8ms)  UPDATE "books" SET "author_id" = NULL WHERE "books"."author_id" = ?  [["author_id", 1]]
=> #<ActiveRecord::Associations::CollectionProxy []>
```
```ruby
class Author < ApplicationRecord 
  has_many :books, dependent: :destroy
end 
```
```ruby
irb(main):002:0> a.books.clear
  SQL (0.4ms)  DELETE FROM "books" WHERE "books"."author_id" = ?  [["author_id", 1]]
=> #<ActiveRecord::Associations::CollectionProxy []>
```

```ruby
class Author < ApplicationRecord
  has_many :books, dependent: :delete_all
end 
```
```ruby
irb(main):002:0> a.books.clear
  SQL (0.2ms)  DELETE FROM "books" WHERE "books"."author_id" = ?  [["author_id", 1]]
=> #<ActiveRecord::Associations::CollectionProxy []>
```
- `books.empty?`
- `books.size`
- `books.find(...)`
- `books.where(...)`



