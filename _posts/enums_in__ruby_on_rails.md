
---
label: til
layout: default
title: enums in Ruby on rails
---
Model

```ruby
class Category < ApplicationRecord
enum category_type: {News: 0, Recruitment: 1}
end
```
Migration

```ruby
class AddCategoryTypeToCategories < ActiveRecord::Migration[5.0]
  def change
    add_column :categories, :category_type, :integer, default: 0
  end
end
```

How to use

```ruby
Category.category_types
#=> ["News", "Recruitment"]
Category.first.category_type
#=> "News"
Category.first.category_type_before_type_cast
#=> 0
```

