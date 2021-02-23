
---
label: til
layout: default
title:  migrate in Rails
---
### add index to column in by migrate
```yml
rails g migration add_part_number_to_products part_number:string:index
```
### references
```yml
bin/rails generate migration AddUserRefToProducts user:references 
```
generates
```ruby
class AddUserRefToProducts < ActiveRecord::Migration[5.0]
  def change
    add_reference :products, :user, foreign_key: true
  end
end
```
This is migration will create a `user_id` column and appropriate index.

### add index
1. `t.string :name, index: true`
2. `add_index :table_name, :column_name`
3. in `create_join_table`:
`t.index `
  

