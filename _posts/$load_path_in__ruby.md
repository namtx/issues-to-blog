
---
label: til
layout: default
title: $LOAD_PATH in Ruby
---
When you add `/Users/namtx/workspace/lightning-pocker/lib` to $LOAD_PATH, you can use
```ruby
require 'card'
```
Instead of 
```ruby
require '/Users/namtx/workspace/lightning-pocker/lib/card'
```
#### $LOAD_PATH

```bash
namtx@Trans-MacBook-Pro ~/w/n/r/lightning-pocker> irb
2.3.0 :001 > puts $LOAD_PATH
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/gems/2.3.0/gems/did_you_mean-1.0.0/lib
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/site_ruby/2.3.0
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/site_ruby/2.3.0/x86_64-darwin18
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/site_ruby
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/vendor_ruby/2.3.0
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/vendor_ruby/2.3.0/x86_64-darwin18
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/vendor_ruby
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0
/Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0/x86_64-darwin18
 => nil
```

#### Usage
And to ensure `lib` directory is load first:
```ruby
$LOAD_PATH.unshift File.expand_path('../lib', __FILE__)
puts $LOAD_PATH

# /Users/namtx/workspace/nimble-university/rspec-ruby-application-testing/lightning-pocker/lib
# /Users/namtx/.rvm/rubies/ruby-2.3.0/lib/ruby/gems/2.3.0/gems/did_you_mean-1.0.0/lib
# ...
```

