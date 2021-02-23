
---
label: til
layout: default
title: Callbacks in Ruby
---
### Use block 
```ruby
def x num 
  yield num
end 
x 3 {|y| y*y} #=> 9
```
Or perhaps converted to a `Proc`; here I show that the "block", converted to a `Proc` implicitly with &block, is just another "callable" value:
```ruby
def x z, &block
  callback = block
  callback.call z
end

# look familiar?
x 4 {|y| y * y} # => 16 
```
### Or lambda
```ruby
def x z,fn
  fn.call z
end

# just use a lambda (closure)
x 5, lambda {|y| y * y} # => 25
```
### Methods 
While the above approaches can all wrap "calling a method" as they create closures, bound `Methods` can also be treated as first-class callable objects:
```ruby
class A
  def b(z)
    z*z
  end
end

callable = A.new.method(:b)
callable.call(6) # => 36

# and since it's just a value...
def x(z,fn)
  fn.call(z)
end
x(7, callable) # => 49 
```
In addition, sometimes it's useful to use the #send method (in particular if a method is known by name). Here it saves an intermediate Method object that was created in the last example; Ruby is a message-passing system:

```ruby
# Using A from previous
def x(z, a):
  a.__send__(:b, z)
end
x(8, A.new) # => 64
```


