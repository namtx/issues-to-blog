
---
label: til
layout: default
title: References in Rust 
---
![image](https://user-images.githubusercontent.com/25602820/28834793-888f8c20-770d-11e7-8cc7-e911b1647e08.png)
```rust
fn main() {
    let s = String::from("hello");
    let len = calculate_length(&s);
    println!("The length of '{}' is: {}", s, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
```
Change value references 

```rust
fn main() {
    let s = String::from("hello");
    change(&s);
}

fn change(some_string: &String) -> String {
    some_string.push_str(", world!");
}
// Doesn't work!
```
#### Mutable References 

```rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world!");
}
// It works like a charm!
```

