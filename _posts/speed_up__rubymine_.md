
---
label: til
layout: default
title: Speed up Rubymine 
---
```ruby
# ~/Library/Preferences/RubyMine2016.<VERSION>/rubymine.vmoptions

-Xms2g
-Xmx4g
-XX:ReservedCodeCacheSize=240m
-XX:+UseCompressedOops

# Restart RubyMine
```

