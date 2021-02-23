
---
label: til
layout: default
title: Immutable fromJS 
---
https://stackoverflow.com/a/34613607/3898402
In this example,

```var a = {address: {postcode: 5085}}
var d = Immutable.Map(a);
```
Here, `d.get('address')` is immutable. It's value cannot change to any other objects. We can only create a new Object from the existing object using the `Immutable.Map.set()` function of `ImmutableJS`.

But, the object referenced by `d.get('address') i.e, {postcode:5085}` is a standard JavaScript object. It is mutable. A statement like this can alter the value of postcode :

```d.get('address').postcode=6000;```
If you check the value of d again, you can see that he value has been changed.

```console.log(JSON.stringify(d));   //Outputs {"address":{"postcode":6000}}```
which is against the principles of immutability.

The reason is that ImmutableJS data structures like List and Map imparts the immutability feature to only the level-1 members of the List/Map.

So, if you have objects inside arrays or arrays inside objects and want them too to be immutable, your choice is Immutable.fromJS.
```
var a = {address: {postcode: 5085}}
var b = Immutable.fromJS(a);
b.get('address').postcode=6000;
console.log(JSON.stringify(b));   //Outputs {"address":{"postcode":5085}}
```
From the above example you can clearly know how fromJS makes the nested members immutable.

I hope you understood the difference between Map and FromJS. All the best =)

