
---
label: til
layout: default
title: Stream and stdin, stdout, stderr
---
- Create a new file & feed some data
```bash
$ touch /tmp/til.txt
$ printf "Look, I have something inside" > /tmp/til.txt
```
- Now
```bash
$ cat /tmp/til.txt                                                                                                        
Look, I have something inside% 
```
- Open a file for reading
```bash
$ exec 3< /tmp/til.txt
```
- Now, you can read 
```bash
$ cat 0<&3
Look, I have something inside%
```

> Once you have opened a file descriptor for reading, you can only “consume” it once
> 

- Run again
```bash
$ cat 0<&3

```
nothing shown
- We also can create a file-descriptor to write file
```bash
$ touch /tmp/til-2.txt
$ exec 4>&/tmp/til-2.txt
```
- And then,
```bash
$ cowsay Hello there! 1>&4
 ______________
< Hello there! >
 --------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

