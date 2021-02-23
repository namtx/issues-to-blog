
---
label: til
layout: default
title: Caching in Rails 
---
## Page Caching 
The whole `HTML` page is saved to a file inside the `public` directory. On subsequent requests, this file is being sent directly to the user without the need to render the view and layout again.

_Gemfile_
```ruby 
[...]
gem 'actionpack-page_caching'
[...]
```
_config/application.rb_
```ruby
[...]
config.action_controller.page_cache_directory = "#{Rails.root.to_s}/public/deploy"
[...]
```
_pages_controller.rb_
```ruby
class PagesController < ApplicationController
  def index
  end
end
```
_config/routes.rb_
```ruby
[...]
root 'pages#index'
[...]
```
_views/pages/index.html.erb_
```erb
<h1>Welcome to my Cached Site!</h1>
```
_pages_controller.rb_
```ruby
class PagesController < ApplicationController
  caches_page :index
end
```
Boot the server and navigate to the root path. You should see the following output in the console:
```ruby
Write page f:/rails/my/sitepoint/sitepoint-source/BraveCacher/public/deploy/index.html (1.0ms)
```
Expire caching 
_pages_controller.rb_
```ruby
expire_page action: 'index'
```

