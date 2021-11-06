# Table: hypothesis_profile

If you [authenticate](../index.md), this table reports your Hypothesis username, display name, authority, and groups.

## Examples

### Get your username, display name, and authority

```
select
  "user",
  display_name,
  authority
from
  hypothesis_profile
```

### Get the names and ids of groups you belong to

```
select
  jsonb_array_elements(groups)
from
  hypothesis_profile
```

### Among the most recent 1000 annotions, find those in your private groups

```
with groups as (
  select
    jsonb_array_elements(groups) as group_info
  from
    hypothesis_profile
),
annos as (
  select
    *
  from
    hypothesis_search
  where 
    query = 'limit=1000'
  order by
    created desc
)
select
  'https://hypothes.is/a/' || a.id as link,
  a."group",
  g.group_info ->> 'name' as name,
  a."user",
  a.created,
  a.title,
  a.uri
from 
  groups g
join
  annos a
on 
  g.group_info ->> 'id' = a."group"
where
  g.group_info ->> 'public' != 'true'
```    



