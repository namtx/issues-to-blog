
---
label: til
layout: default
title: Code Smell
---
### Long method
 + flog


 `Extract Method`

 `Replace Temp with Query`
### Large Class
difficult to understand and harder to change or reuse => more bugs and conflicts
#### Symptoms
+ Can't describe what the class does in one sentence
+ Can't tell what class does without scrolling
+ Class needs to change for more than one reason
+ more than 7 methods
+ total `flog` score of 50

#### Solutions
+ Move method
+ Extract Class
+ Replace Conditional with Polymorphism
+ Extract Value Object
+ Extract Decorator
+ Replace Subclasses with Strategies

#### Prevention
+ Single Responsibility Principle
+ Composition Over Inheritance
+ Dependency Inversion Principle
+ Open/Closed Principle

#### Private Method
+ Private methods can't be reused between classes => duplicated => extract private methods to new classes
+ Private methods can't be tested directly
+ Private methods are often the easiest to extract to new classes
+ Keeping a class's private interface as small as possible is a best practice as well.

#### God Class
God class is class that seems to know everything about an application

`User` class is a sample

### Feature Envy
Feature envy reveals a method that would work better on a different class

#### Symptoms
+ Repeated references to the same object
+ Parameters or local variables which are used more than methods and instance variables of the class in question
+ Method that includes a class name in their own names (`invite_user`)
+ Private methods on the same class that accept the same parameter
+ `Law of Demeter`
+ `Tell, Don't Ask`

#### Solutions
+ Extract method
+ Move method
+ In-line class

#### Prevention
+ `Tell, don't ask`
+ `Law of Demeter`

### Case Statement
Sign that a method contains too much knowledge

#### Symptoms
+ Check the class of the class of an object
+ check a type code
+ adding `when` clauses

#### Solutions
+ Replace Type Code with Subclasses
+ Replace Conditional with Polymorphism
+ Use Convention over Configuration

### Long Parameter List

#### Solutions
+ Introduce Parameter Object
+ Extract Class

#### Anti-Solution
+ Use hash of named parameters

### Duplicated Code
+ Don't repeat yourself

#### Symptoms
+ Copy and pasting
+ Shotgun Surgery

#### Solutions
+ Extract Method
+ Extract Class
+ Extract Partial
+ Replace Conditional with Polymorphism
+ Replace Conditional with Null Object

  Remove duplicated checks for `nil` values

#### Prevention
+ Follow Single Responsibility Principle

### Uncommunicative Name

#### Symptoms
+ Difficult understanding a method or class
+ Methods or classes with similar names but dissimilar functionality
+ Redundant names, such as names which include type of object to which they refer.

#### Solutions
+ Rename method
+ Extract class
+ Extract method
+ In-line class

### STI

#### Symptoms
+ You need to change from one subclass to another
+ Behavior is shared among some subclass but not others.
+ One subclass is a fusion of one or more other subclasses.

#### Solutions
+ If you're using STI to reuse common behavior, use `Replace Subclasses with Strategies` to switch to a composition-based model.
+ If you're using STI so that you can easily refer to several different classes in the same table, switch to using a polymorphic association instead.

#### Prevention
+ Composition over Inheritance

### Comments

### Mixin

### Callback


