
---
label: til
layout: default
title: Multiple export with dot 
---
### Do you want to use `Sidebar.Item` instead of `SidebarItem`?

```javascript
import Siderbar from './sidebar';
import Item from './sidebar-item';

export default Object.assign(Sidebar, { Item });
```

