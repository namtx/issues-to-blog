
---
label: til
layout: default
title: super keyword in Ruby's inheritance
---
```ruby
class Animal 
  def initialize name
    @name = name 
  end
end

class Dog 
  def initialize color
    super
    @color = color
  end
end

bruno = Dog.new "brown"
p bruno #=> < @name = "brown", @color = "brown" >
```
in addition to the default behavior, `super` automatically forwards the arguments that were passed to the method from which `super` is called (`initialize` method in `Dog` class)

Right way!

```ruby
...
class Dog
  def initialize name, color
    super #=> wrong number of arguments
    super name
    @color = color
  end
end

bruno = Dog.new "bruno", "brown"
p bruno #=> < @name = "bruno", @color = "brown" >
```


