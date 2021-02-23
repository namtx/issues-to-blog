
---
label: til
layout: default
title: JAVA_HOME and libexec/java_home
---
### Get the list of available JDKs
```bash
$ /usr/libexec/java_home -V
Matching Java Virtual Machines (5):
    15.0.1, x86_64:     "OpenJDK 15.0.1"        /Users/namtx/Library/Java/JavaVirtualMachines/openjdk-15.0.1/Contents/Home
    13.0.5, x86_64:     "Zulu 13.35.17" /Users/namtx/Library/Java/JavaVirtualMachines/azul-13.0.5/Contents/Home
    11.0.9.1, x86_64:   "Zulu 11.43.55" /Users/namtx/Library/Java/JavaVirtualMachines/azul-11.0.9.1/Contents/Home
    1.8.0_275, x86_64:  "AdoptOpenJDK 8"        /Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home
    1.8.0_272, x86_64:  "Zulu 8.50.0.21"        /Users/namtx/Library/Java/JavaVirtualMachines/azul-1.8.0_272/Contents/Home

/Users/namtx/Library/Java/JavaVirtualMachines/openjdk-15.0.1/Contents/Home
```
### Specific JDK version by libexec/java_home
```bash
$ export JAVA_HOME=$(/usr/libexec/java_home -v 11.0.9.1)

$ java -version
openjdk version "11.0.9.1" 2020-11-04 LTS
OpenJDK Runtime Environment Zulu11.43+55-CA (build 11.0.9.1+1-LTS)
OpenJDK 64-Bit Server VM Zulu11.43+55-CA (build 11.0.9.1+1-LTS, mixed mode)

```

