
---
label: til
layout: default
title: best practices for Presenter Rails 
---
1. _Your controller action should ideally expose only one instance variable to your views - the presenter._
2. _Back each view or partial with its own presenter._
3. _Rendering partials within partials means your presenters can return other presenters for these partials._
4. _If your view requires an author and the current user, provide that to your presenter - don't make your presenter navigate through an item to get the author._
5. _Write your view first, assuming no knowledge of how our domain is modelled, then write the presenter to drive it._
6. _Prefer composition over inheritance._
7. _Presenters shouldn't know about CSS class names._

