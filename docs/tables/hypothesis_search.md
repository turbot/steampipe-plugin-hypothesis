# Table: hypothesis_search

Searches for Hypothesis annotations matching a query.

## Examples

### Search for recent annotations by a person

**NOTE** `group` and `user` are special to Postgres so, for those columns only, you have to double-quote the names.

```sql
select
  created,
  url,
  text,
  tags
from
  hypothesis_search
where
  query = 'user=judell';
```


