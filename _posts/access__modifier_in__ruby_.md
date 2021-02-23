
---
label: til
layout: default
title: Access Modifier in Ruby 
---
+ `public`:  phương thức này có thể gọi từ bất cứ đâu, có thể coi như không có kiểm soát hay giới hạn truy cập.
+ `protected`: phương thức này có thể được gọi từ một đối tượng của lớp này hoặc một đối tượng là con của lớp này.
+ `private`: phương thức này chỉ được gọi từ chính nó (đối tượng gọi và đối tượng nhận là một).
### Truy cập từ một đối tượng khác lớp (đối tượng gọi phương thức và đối tượng nhật phương thức thuộc 2 lớp khác nhau)
```ruby
class AccessControll
  def test_public
    puts "public access OK"
  end

  protected

  def test_protected
    puts "protected access OK"
  end

  private

  def test_private
    puts "private access OK"
  end
end

class Test
  def test
    aceess_controll = AccessControll.new
    aceess_controll.test_public # public access OK
    aceess_controll.test_protected # NoMethodError
    aceess_controll.test_private # NoMethodError
  end
end
Test.new.test
```
### Truy cập từ một đối tượng khác thuộc cùng một lớp (2 đối tượng gọi và nhận phương thức thuộc cùng một lớp)
```ruby
class AccessControll
  def test other
    other.test_public # public access OK
    other.test_protected # protected access OK
    other.test_private # NoMethodError
  end

  def test_public
    puts "public access OK"
  end

  protected

  def test_protected
    puts "protected access OK"
  end

  private

  def test_private
    puts "private access OK"
  end
end

AccessControll.new.test AccessControll.new 
```

### Truy cập từ một đối tượng thuộc lớp con (đối tượng gọi có lớp là lớp con của đối tượng nhận)
```ruby
class AccessControll
  def test_public
    puts "public access OK"
  end

  protected

  def test_protected
    puts "protected access OK"
  end

  private

  def test_private
    puts "private access OK"
  end
end

class Children < AccessControll
  def test other
    other.test_public # public access OK
    other.test_protected # protected access OK
    other.test_private # NoMethodError
  end
end

Children.new.test AccessControll.new
```

### Truy cập từ một đối tượng thuộc lớp cha (đối tượng gọi có lớp là lớp cha của đối tượng nhận)

```ruby
class Parent
  def test other
    other.test_public # public access OK
    other.test_protected # protected access OK
    other.test_private # NoMethodError
  end
end

class AccessControll < Parent
  def test_public
    puts "public access OK"
  end

  protected

  def test_protected
    puts "protected access OK"
  end

  private

  def test_private
    puts "private access OK"
  end
end

Parent.new.test AccessControll.new 
```

### Truy cập từ cùng một đối tượng (đối tượng gọi phương thức và đối tượng nhận phương thức là 1 đối tượng duy nhất)
```ruby
class AccessControll
  def test
    test_public # public access OK
    test_protected # protected access OK
    test_private # private access OK
  end

  def test_public
    puts "public access OK"
  end

  protected

  def test_protected
    puts "protected access OK"
  end

  private

  def test_private
    puts "private access OK"
  end
end
AccessControll.new.test
```

