
---
label: til
layout: default
title: belongs_to association in Rails
---
## Methods
```ruby
class Book < ApplicationRecord
  belongs_to :author
end
```
Book class will gain 5 methods related to association
```ruby
author
author=(association)
build_author(attributes: {})
create_author(attributes: {})
create_author!(attributes: {})
```

## Options
### 1. auto_save
### 2. class_name
### 3. counter_cache
```ruby
class Author < ApplicationRecord
  has_many :books 
end 

class Book < ApplicationRecord
  belongs_to :author, counter_cache: true
end 
```
when asking value for `@author.books.size`, `counter_cache` avoid perform query `COUNT(*)` to database multiple times by save `size` value to cache. 

actual column must be added to the associated (`has_many`) model. You would need to add column `books_count` to the Author model.
### 4. dependent
- `:destroy` associated objects to also be destroyed
- `:delete_all` associated objects to also be deleted from database (callbacks are not executed)
- `:nullify` foreign_key keys to be set NULL (callbacks are not executed)
- `:restricted_with_exception` causes an exception to be raised if there are associated records
- `:restricted_with_error` causes an error to the owner if there are associated objects
### 5. foreign_key
lets you set the name of `foreign_key` directly.
### 6. inverse_of
### 7. polymorphic
### 8. touch
if you set option `touch: true`, whenever associated be added or destroyed, `updated_at` or `updated_on` will be set to current time
### 9. validate
if you set option `validate: true`, whenever object is save, then associated object will be validated.
### 10. optional
if option `optional: true`. the presence of associated object will not be validated. 
## Scopes

### 1. where
let you specify the conditions that associated object must meet
```ruby
class Book < ApplicationRecord
  belongs_to :author, ->{where active: true}, dependent: :destroy
end
```
### 2. includes
```ruby

class LineItem < ApplicationRecord
  belongs_to :book
end
 
class Book < ApplicationRecord
  belongs_to :author
  has_many :line_items
end
 
class Author < ApplicationRecord
  has_many :books
end
```
If you frequently retrieve directly authors from line items (`@line_items.book.author`), you can make your code more efficiently by  add `includes` 
```ruby
class LineItem < ApplicationRecord
  belongs_to :book, -> { includes :author }
end
 
class Book < ApplicationRecord
  belongs_to :author
  has_many :line_items
end
 
class Author < ApplicationRecord
  has_many :books
end
```
### 3. readonly
### 4. select
## Assign any object to the `belongs_to` association doesn't automatically save the object.

