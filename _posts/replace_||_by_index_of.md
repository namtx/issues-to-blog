
---
label: til
layout: default
title: Replace || by indexOf
---
```js 
if (input.files && input.files[0] && (ext == 'gif' || ext == 'png' || ext == 'jpeg' || ext == 'jpg'))...
```

=> 

```js
if (input.files && input.files[0] && ['gif', 'png', 'jpeg', 'jpg'].indexOf(ext) > -1)
```

