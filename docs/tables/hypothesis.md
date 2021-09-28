# Table: hypothesis

This table displays a hardcoded greeting.

## Examples

```
> select *, pg_typeof(json) from hypothesis
+----+----------+-------------------+-----------+
| id | greeting | json              | pg_typeof |
+----+----------+-------------------+-----------+
| 1  | hypothesis    | {"hypothesis":"world"} | jsonb     |
| 2  | hypothesis    | {"hypothesis":"world"} | jsonb     |
| 3  | hypothesis    | {"hypothesis":"world"} | jsonb     |
+----+----------+-------------------+-----------+

> select *, json->>'hypothesis' as json_value from hypothesis where id = 2
+----+----------+-------------------+------------+
| id | greeting | json              | json_value |
+----+----------+-------------------+------------+
| 2  | hypothesis    | {"hypothesis":"world"} | world      |
+----+----------+-------------------+------------+
```

