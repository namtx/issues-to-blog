
---
label: til
layout: default
title: associations in Rails
---
## self joins
```ruby
class Employee < ApplicationRecord
  has_many :subordinates, class_name: Employee.name, foreign_key: :manager_id
  belongs_to :manager, class_name: Employee.name
end
```
> Migration

```ruby
class CreateEmployees < ActiveRecord::Migration[5.0]
  def change
    create_table :employees do |t|
      t.references :manager, foreign_key: true
      t.string :name

      t.timestamps
    end
  end
end
```

## controlling caching
Rails keeps result of most recent query available for further operation. 
```ruby
author.books # query
author.books.size # cache
author.books.empty? # cache
```
Use `reload` function to reload cache
## name collision
Creating a association adds a new method to model, It's a bad idea to give association a name that already used for instance methods of `ActiveRecord::Base`.
Association method would override the base method and breaks thing. 
## updating schema
```ruby
class Author < ApplicationRecord
  has_many :books
end 

class Book < ApplicationRecord
  belongs_to :author
end
```
You `need` to backed up by proper foreign key declaration on then books table:
```ruby
class CreateBooks < ActiveRecord::Migration[5.0]
  def change
    create_table :books do |t|
      t.datetime :published_at
      t.string   :book_number
      t.integer  :author_id
    end
  end
end
```

It's good practice to add an index on the foreign key to improve queries performance 
```ruby
class CreateBooks < ActiveRecord::Migration[5.0]
  def change
    create_table :books do |t|
      t.datetime :published_at
      t.string   :book_number
      t.integer  :author_id
    end
 
    add_index :books, :author_id
    add_foreign_key :books, :authors
  end
end
```
### `inverse_of`
Normal association: 
```ruby
class Author < ApplicationRecord
  has_many :books
end

class Book < ApplicationRecord
  belongs_to :author
end
```
`bi-directional` association
Active Record will load only one copy of Author object, makes application more efficient
```ruby
a = Author.first
b = a.books.first
a.equal? b.author #true
```
However, Active Record will not automatically identify `bi-directional` associations that contain any of following options: 
- conditions
- class_name
- through
- polymorphic
- foreign_key

```ruby
a = Author.first
b = a.books.first
a.equal? b.author #false
```
Active Record provides `:inverse_of` keyword, you will explicitly declare `bi-directional` association 
```ruby
class Author < ApplicationRecord
  has_many :books, inverse_of: :writer
end 

class Book < ApplicationRecord
  belongs_to :writer, class_name: Author.name, foreign_key: :author_id
end
```
```ruby
a = Author.first
b = a.books.first
a.equal? b.writer # true
```











