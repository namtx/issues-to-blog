
---
label: til
layout: default
title: Export css module variable to JS
---
```scss
// styles.module.scss
:export {
  grayLightColor: $color-gray-light;
}
```

```js
import styles from './styles.module.scss';

console.log(styles.grayLightColor);
```

