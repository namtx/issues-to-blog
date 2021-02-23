
---
label: til
layout: default
title: Docker log with lnav
---
### Redirect Docker Logs to File
```bash
docker log -f [container_id] > docker.log
```

### lnav log file
```bash
lnav docker.log
```

