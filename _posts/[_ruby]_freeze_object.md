
---
label: til
layout: default
title: [Ruby] freeze object
---
## Creating immutable objects

> In Ruby constant is mutable

```ruby
MY_CONSTANT = "foo"
MY_CONSTANT << "bar"
p MY_CONSTANT
#=> "foo bar"
```
But with freeze, when you attempt to change the frozen object -> Runtime Error
```ruby
MY_CONSTANT = "foo".freeze
MY_CONSTANT << " bar"
#=> `<main>': can't modify frozen String (RuntimeError)
```
Real-world example
ActionDispatch codebase, Rails hides sensitive data by logging '[FILTERED]' literal string, this text is stored in constant frozen.
```ruby
module ActionDispatch
  module Http
    class ParametersFilter
      FILTER = '[FILTERED]'.freeze
...
```
## Reducing object allocation
One of best things to speed up Ruby on decrease number of objects are created.
In your app, a method call `log("foobar")`, it is creating a String object.
With freeze, Ruby interpreter will create only one String literal and caches it for feature use. 
> 50% performance increase
```ruby
require "benchmark/ips"

def noop arg
end

Benchmark.ips do |x|
  x.report("normal"){noop "foo"}
  x.report("frozen"){noop "foo".free}
end
```

> Warming up --------------------------------------
              normal   328.890k i/100ms
              frozen   357.420k i/100ms
Calculating -------------------------------------
              normal      8.535M (± 2.4%) i/s -     42.756M in   5.012534s
              frozen     10.428M (± 4.4%) i/s -     52.183M in   5.014607s

Real-world example, in Rails routes, the router is used for every request => freeze String literal.
```ruby
# excerpted from https://github.com/rails/rails/blob/f91439d848b305a9d8f83c10905e5012180ffa28/actionpack/lib/action_dispatch/journey/router/utils.rb#L15
def self.normalize_path(path)
  path = "/#{path}"
  path.squeeze!('/'.freeze)
  path.sub!(%r{/+\Z}, ''.freeze)
  path.gsub!(/(%[a-f0-9]{2})/) { $1.upcase }
  path = '/' if path == ''.freeze
  path
end
```
## Ruby >= 2.2
Optimization is included in `Ruby 2.2` and later.
frozen string literal is used for hash key
```ruby
user["name"] #Ruby 2.2
user["name".freeze] #Ruby 2.1
```
## Freeze object in `initialize` methods

