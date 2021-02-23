
---
label: til
layout: default
title: Ruby method lookup path 
---
1. Lookup within singleton class
```ruby
class MyCar
  def method
    "defined on car"
  end
end

object = MyCar.new

def object.method
  "defined on singleton"
end

object.method # => "defined on singleton"
```
2. Look within modules that extended singleton 
```ruby 
module MyModule
  def method; 'defined on MyModule'; end
end

object = MyCar.new
object.extend(MyModule)
object.method # => 'defined on MyModule'
```
If you extend with multiple modules, later modules take precedence:
```ruby
module MyModule2
  def method; 'defined on MyModule2'; end
end

object.extend(MyModule)
object.extend(MyModule2)
object.method # => 'defined on MyModule2' 
```
3. Look within methods defined on class 
```ruby 
class MyCar
  def method; 'defined on MyCar'; end
end

MyCar.new.method # => 'defined on MyCar'
```
4. Look within modules that were mixed in when class was defined 
```ruby
module A
  def method; 'defined on A'; end
end

module B
  def method; 'defined on B'; end
end

class MyCar
  include A
  include B
end 
```
Similar to when we extend a singleton, the later modules take precedence:
```ruby
MyCar.new.method # => 'defined on B' 
```
5. Climb the ancestors chain
=> move up ancestors chain and start from (3)
6. Start again, check method_missing
We've reached the end of the ancestors chain and still find out method. Ruby will return back (1) and run `method_missing`
```ruby
class BasicObject
  def method; 'defined on Object'; end
end

class MyCar
  def method_missing(name, *args, &block)
    "called #{name}"
  end
end

MyCar.ancestors  # => [MyCar, Object, Kernel, BasicObject] 
MyCar.new.method # => 'defined on Object'
```
### Super 
The super method looks up the ancestors chain for the method name where super is called. It accepts optional arguments which are then passed onto that method. This will follow the same lookup path as outlined above from when the super method was called.
```ruby
class Vehicle
  def method; 'defined on Vehicle'; end
end

class MyCar < Vehicle
  def method
    "defined on MyCar\n"
    super
  end
end

MyCar.new.method
# => "defined on MyCar"
# => "defined on Vehicle"
```

