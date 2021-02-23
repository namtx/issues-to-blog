
---
label: til
layout: default
title: singleon method & singleton class
---
## Singleton method 
is method is defined for only one instance
```ruby
class Foo
  def bar
    :bar
  end
end

foo = Foo.new
p foo.bar #=> bar

def foo.baz
  :baz
end

p foo.baz #=> baz
```
## Singleton class
is class have only one instance
```ruby
require "singleton" 

class Foo
  include Singleton
end

foo, foo1 = Foo.instance, Foo.instance

foo == foo1 #=> true

Foo.new #=> Runtime Error
```

