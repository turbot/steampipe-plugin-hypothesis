# Table: hypothesis_search

Searches for Hypothesis annotations matching a query. If you [authenticate](../index.md) you'll search the Hypothesis public layer plus all your private annotations, and annotations in private groups you belong to. If you don't authenticate you'll just search the public layer.

## Examples

### Find 10 recent notes, by `judell`, that have tags

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
order by
  created desc
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

### Find notes on the New York Times home page, by month

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
### Find URLs and note counts on articles annotated in the Times' Opinion section

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

### Find notes, in the Times' Opinion section, that quote selections matching "covid"

**NOTE** You do not want to write this kind of SQL more than once. Unpacking Postgres JSONB is tricky. Instead, package JSONB idioms in functions like [quote_from_anno](https://jonudell.info/h/analytics/doc/functions.html#quote_from_anno). See [Postgres functional style](https://blog.jonudell.net/2021/08/21/postgres-functional-style/) for details.

```sql
with expanded_target_keys as (
  select
    jsonb_object_keys(target->0) as target_key,
    *
  from
    hypothesis_search
  where
    query = 'wildcard_uri=https://www.nytimes.com/*/opinion/*'
),
to_selectors as (
  select 
    target->0->'selector' as selectors,
    *
from 
  expanded_target_keys
where  
  target_key = 'selector'
  and target->0->>'selector' is not null
),
to_selector_objects as (
  select
    jsonb_array_elements(target->0->'selector') as selector_object,
    *
from 
  to_selectors
),
to_text_quote_selectors as (
  select
    *
  from
    to_selector_objects
  where
    selector_object->>'type' = 'TextQuoteSelector'
)
select
  'https://hypothes.is/a/' || id as link,
  uri,
  "user",
  created,
  selector_object->>'exact' as quote
from 
  to_text_quote_selectors
where
  selector_object->>'exact' ~* 'covid'
order by
  created desc
```  

### Find annotated GitHub repos, join with info from GitHub API

**NOTE** This will take a minute or so. Once it's done, it's cached for 5 minutes, or another duration you can specify, so queries that touch the same data are instantaneous.

```sql
with annotated_urls as (
  select
    regexp_matches(uri, 'github.com/([^/]+)/([^/]+)') as match,
    *
  from 
    hypothesis_search
  where
    query = 'wildcard_uri=http://github.com/*&limit=1000'
  order by
    uri
  ),
and_repos as (
  select
    *,
    match[1] || '/' || match[2] as repository_full_name
from 
  annotated_urls
order by 
  uri
)
select distinct
  g.name,
  g.description,
  g.owner_login,
  r.uri,
  r.id,
  r."user"
from 
  github_repository g
join 
  and_repos r
on
  g.full_name = r.repository_full_name
```

### Order annotated GitHub repos by count of annotations

```sql
with annotated_urls as (
  select
    id,
    created,
    uri,
    "user",
    regexp_matches(uri, 'github.com/([^/]+)/([^/]+)') as match    
  from 
    hypothesis_search
  where
    query = 'wildcard_uri=http://github.com/*&limit=1000'
  order by
    uri
  ),
and_repos as (
  select
    a.*,
    a.match[1] || '/' || a.match[2] as repository_full_name
from 
  annotated_urls a
order by 
  a.uri
),
joined as (
  select distinct
    g.name,
    g.owner_login,
    r.uri,
    r.id,
    r."user"
  from 
    github_repository g
  join 
    and_repos r
  on
    g.full_name = r.repository_full_name
)
select 
  count(j.*),
  j.name,
  j.owner_login,
  j.uri
from
  joined j
group by
  j.name,
  j.owner_login,
  j.uri
order by
  count desc
```

