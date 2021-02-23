
---
label: til
layout: default
title: ancestors in Ruby
---
```ruby
module Speak
  def speak sound
    puts "#{sound}"
  end
end 

class Dog
  include Speak
end

class Human 
  include Speak
end

puts "---Dog ancestors---"
puts Dog.ancestors
puts ''
puts "---Human ancestors---"
puts Human.ancestors
```
Result:
```sh
---Dog ancestors---
Dog
Speak
Object
Kernel
BasicObject

---Human ancestors---
Human
Speak
Object
Kernel
BasicObject
```
When function is called on `Dog` or `Human` instance, it will look up on it's `class` first and then it's `module`

