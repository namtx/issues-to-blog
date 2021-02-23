
---
label: til
layout: default
title: Install PostgresQL on Ubuntu 16.04
---
- Install 
```bash
    $ sudo apt-get update
    $ sudo apt-get install postgresql postgresql-contrib
```
- Switch to `postgres` user
```bash
   $ sudo -i -u postgres
```
- Create `root` account
```bash
    createuser --interactive
```
Input `root`, as superuser
- Create `root` database for user `root`
```
    createdb root
```
- Alter `root`'s password
```
    psql
    alter user root with encrypted password 'toor'
```

Now you can access to PostgreSQL by `root` and password `toor`



