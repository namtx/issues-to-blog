
---
label: til
layout: default
title: MySQL datatypes 
---
# Schema Optimization and Indexing

###  Whole numbers
  ```sql
    TINYINT -- 8 bits
    SMALLINT -- 16 bits
    MEDIUMINT -- 24 bits
    INT -- 32 bits
    BIGINT -- 64 bits
  ```
  Store value from `- 2 ** (N-1)` to `2 ** (N-1) - 1`
  - `SIGNED` and `UNSIGNED`: `TINYINT UNSIGNED` store values from 0 - 255 instead of -128 - 127 as `TINYINT`
  - Your choise determines how MySQL store values in memory, disk, but, MySQL use `BIGINT` 64 bits in computations, even on 32 bits architect.
  - MySQL let you define "width", example INT(11), which defines the number of characters MySQL's interactive tools, display purpose, doesn't restrict the range of value: INT(1) and INT(11) are identical.


###  Real numbers
  - `FLOAT`, `DOUBLE` types support approximate calculations with standard floating-point math.
  - `DECIMAL` types is for storing exact fractional numbers.
  - MySQL 5.0 or newer, supports exact math on `DECIMAL`, MySQL 4.1 or earlier used floating-point match to perform calculation on DECIMAL values.
  - MySQL 5.0 or newer itself performs `DECIMAL` exact math, because CPUs don't support compuations directly, floating-point math is faster because CPUs natively supports.

#### Precision
  - `DECIMAL` columns, you can specify maximum allowed digits before and after decimal point -> influences the column's space consumption
  - MySQL 5.0 or newer pack the digits into a binary string (9 digits per 4 bytes): Ex: `DECIMAL(18,9)` will store 9 digits from each side, 4 bytes for 9 digits before decimal point, 1 byte for decimal point, 4 bytes for 9 digits after decimal point.
  - `DECIMAL` number in MySQL 5.0 or newer can have up to 65 digits, earlier MySQL version had a limit 254 bits and store values as unpacked string (one by per digits), use `DOUBLE` for computation.
#### `DECIMAL` use cases
  - Floating-point types typically use less space than `DECIMAL`, `FLOAT` use 4 bytes, `DOUBLE` uses 8 bytes.
  - MySQL uses `DOUBLE` for computations
  - `DECIMAL` require space and computational costs -> use when you need exact results for fractical numbers: ex: storing financial data


