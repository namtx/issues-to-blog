
---
label: til
layout: default
title: differences in Rails
---
## `update_attribute` & `update_attributes`
- `update_attribute` update single column <> `update_attributes`
- `update_attribute` will not perform validations <> `update_attributes`
## `update_columns` & `update_attributes`
- `update_columns`: `callback` is skipped, `validation` is skipped
> + Updates the attributes directly in the database issuing an UPDATE SQL statement
> + This method raises an +ActiveRecord::ActiveRecordError+ when called on new objects, or when at least one of the attributes is marked as readonly
## `update_columns` & `update_column`
- `update_column` is `update_columns` in single column
## `save` without validation
```ruby
save validate: false
```

