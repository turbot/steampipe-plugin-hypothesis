# Table: hypothesis_search

Searches for Hypothesis annotations matching a query. If you [authenticate](../index.md) you'll search the Hypothesis public layer plus all your private annotations, and annotations in private groups you belong to. If you don't authenticate you'll just search the public layer.

## Examples

### Find 10 recent notes, by a person, that have tags

**NOTE** `group` and `user` are special to Postgres so, for those columns only, you have to double-quote the names.

```sql
select
  created,
  uri,
  tags,
  "group", -- always __world__ if unauthenticated
  "user"
from
  hypothesis_search
where
  query = 'user=judell'
  and jsonb_array_length(tags) > 0
limit 10;
```
### Find notes tagged with both `media` and `review`

**NOTE** This matches notes with:
- "media" and "review"
- "social media" and "review"
- "social media" and "peer review"

```sql
select
  uri,
  tags
from
  hypothesis_search
where
  query = 'tag=media&tag=review';
```

### Find notes tagged with `social media` and `peer review`

```sql
select
  uri,
  tags
from
  hypothesis_search
where
  query = 'tag=social+media&tag=peer+review'
```  

### Find notes for the New York Times home page, by month

```sql
with data as (
  select
    substring(created from 1 for 7) as month,
    uri
  from
    hypothesis_search
  where
    query = 'uri=https://www.nytimes.com'
)
select
  month,
  count(*)
from
  data
group by 
  month
order by
  count desc
```  
### Find URLs and note counts for articles annotated in the Times' Opinion section

```sql
with data as (
  select
    uri
  from
    hypothesis_search
  where
    query = 'wildcard_uri=https://www.nytimes.com/*/opinion/*'
)
select
  count(*),
  uri
from 
  data
group by 
  uri
order by 
  count desc
```  

### Find page notes (i.e. notes referring to the URL, not a selection) on www.example.com

```sql
with target_keys_to_rows as (
  select
    id,
    "user",
    created,
    text,
    uri,
    target,
    jsonb_object_keys(target->0) as target_key
  from
    hypothesis_search
  where
    query = 'uri=https://www.example.com'
  group by
    id, "user", created, text, uri, target
)
select
  'https://hypothes.is/a/' || id as link,
  *
from 
  target_keys_to_rows
where 
  target_key = 'selector'
  and target->0->>'selector' is null
order by
  created desc
```

### Find notes, in the Times' Opinion section, on quotes (selections) matching "covid"

```sql
```  

### Find notes on GitHub repos, join with GitHub API

```sql
with urls as (
  select 
    uri
  from 
    hypothesis_search
  where
    query = 'wildcard_uri=http://github.com/*&limit=2000'
  )
select * from urls
```  






