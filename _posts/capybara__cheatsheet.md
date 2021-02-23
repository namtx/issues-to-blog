
---
label: til
layout: default
title: Capybara Cheatsheet
---
Check a checkbox by `id` and `value`:
```ruby 
find(:css, "#services_status_[value='active']").set true
```
Uncheck: 
```ruby
find(:css, "#services_status_[value='active']").set false 
```

