
---
label: til
layout: default
title: Create folder for example coding
---
```shell
for i in {a..z}
do
  if [ ! -d "$i" ]
  then
    mkdir "$i" && cd "$i" && echo "Successed!"
    exit 0
  fi
done

```

