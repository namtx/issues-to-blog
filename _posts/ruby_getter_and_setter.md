
---
label: til
layout: default
title: Ruby getter and setter
---
In `Ruby`, getters and setters are typically defined by using class method `attr_accessor`
```ruby
class Foo 
  attr_accessor :bar,  :baz
end 

foo = Foo.new 
#=> #<Foo:0x007fada99c2d00>
foo.bar #=> nil 
foo.bar = "jelly"
foo.bar #=> "jelly"
```

`attr_accessor` is different than Rails's `attr_accessible` 
`attr_accessor` sets up getters and setters for those instance variables. 
```ruby
def bar
  @bar
end

def bar= value 
  @bar = value
end
...
```
`attr_reader` will define only the getter methods 
`attr_writer` will define only setter methods 

