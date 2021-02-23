
---
label: til
layout: default
title: Describe table in PostgresQL 
---
```sql
SELECT 
   table_name, 
   column_name, 
   data_type,
   is_nullable
FROM 
   information_schema.columns
WHERE 
   table_name = 'compensations';
```

