
---
label: til
layout: default
title: Ruby Constructor 
---
### Responsibilities 
+ allocates space for the object 
+ assigns instance variables their initial values 
+ returns the instance of that class
```ruby
class Foo
  def initialize baz, biz
    @baz, @biz = baz, biz
  end
end 

foo = Foo.new "a", "b"
```
+ `#initialize` is an `instance method` 
+ `.new` is a `class method` 
> `.new` is the Ruby constructor not the instance method `#initialize`

