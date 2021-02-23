
---
label: til
layout: default
title: fetch Hash in Ruby 
---
```ruby
pets = { cat: "Jess" }
pets[:dinosaur]
# => nil

pets = { cat: "Jess" }
pets[:dinosaur] || "They all died :("
# => "They all died :("

pets = { cat: "Jess" }
pets.fetch(:dinosaur, "They all died :(")
# => "They all died :("

pets.fetch(:dinosaur, Dinosaur.raise_from_the_dead!)

pets.fetch(:dinosaur) { Dinosaur.raise_from_the_dead! } 
```

