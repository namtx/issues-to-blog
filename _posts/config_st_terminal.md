
---
label: til
layout: default
title: config st terminal
---
### Step 1 - Clone st repository
```shell
$ git clone https://git.suckless.org/st
```
### Step 2 - Download Xlib files 
You need the Xlib header files to run `make clean build` for `st` compiling
```shell
$ sudo apt-get install libx11-dev libxft-dev libxext-dev
```
### Step 3 - Compile
`cd` into extracted folder from Step 1
and run:
```shell
$ sudo make clean install
```
 ### Step 4 - Apply `scrollback` patch
- Download diff file
```shell
$ curl -o scrollback.patch https://st.suckless.org/patches/scrollback/st-scrollback-20190331-21367a0.diff
```
- apply diff file
```shell
$ git apply scrollback.patch
```
- Edit file config.h for new keybinding
```c
{ShiftMask,  XK_K,  kscrollup,  {.i = -1} },
{ShiftMask,  XK_J,  kscrolldown,  {.i = -1} },
```
- re-compile
```shell
$ sudo make clean install
```
and now you can scroll up or srolldown by `Shift + K` or `Shift + J`

