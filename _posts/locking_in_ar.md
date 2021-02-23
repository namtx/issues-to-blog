
---
label: til
layout: default
title: Locking in AR
---
## Optimistic Locking
Trong loại này, nhiều người dùng có thể truy cập cùng một đối tượng để đọc giá trị của nó, nhưng nếu hai người dùng thực hiện cập nhật thì sẽ phát sinh mâu thuẫn, chỉ có một người sử dụng sẽ thành công và một trong những người khác sẽ không được thực hiện.
```ruby
p1 = Person.find(1)
p2 = Person.find(1)

p1.first_name = "Michael"
p1.save

p2.first_name = "should fail"
p2.save # Raises a ActiveRecord::StaleObjectError 
```
Để tạo Optimistic locking, bạn có thể tạo ra một field lock_version mà bạn muốn đặt khóa và Rails sẽ tự động kiểm tra trước khi cập nhật các đối tượng. Cơ chế của nó là khá đơn giản. Mỗi lần các đối tượng được cập nhật, giá trị lock_version sẽ được tăng lên. Do đó, nếu hai yêu cầu muốn thực hiện cùng một đối tượng, yêu cầu đầu tiên sẽ thành công vì lock_version của nó cũng giống như khi nó được đọc. Nhưng yêu cầu thứ hai sẽ thất bại vì lock_version đã được tăng lên trong cơ sở dữ liệu của các yêu cầu đầu tiên.
## Pessimistic Locking
Với loại locking này, chỉ có người dùng đầu tiên truy cập đến các đối tượng sẽ có thể cập nhật nó. Tất cả những người dùng khác sẽ bị loại khỏi thậm chí đọc các đối tượng (hãy nhớ rằng trong Optimistic locking, chúng tôi chỉ khóa nó khi cập nhật dữ liệu và người dùng vẫn có thể đọc nó).

Rails sẽ thực hiện Pessimistic Locking bằng cách phát hành truy vấn đặc biệt trong cơ sở dữ liệu. Ví dụ, giả sử bạn muốn lấy đối tượng tài khoản và khóa nó cho đến khi bạn hoàn thành việc cập nhật:
```ruby
account = Account.find_by_user_id(5)
account.lock!
#no other users can read this account, they have to wait until the lock is released
account.save! 
#lock is released, other users can read this account
```

