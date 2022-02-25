# Table: hypothesis_search

Searches for Hypothesis annotations matching a query. If you [authenticate](https://hub.steampipe.io/plugins/turbot/hypothesis#credentials) you'll search the Hypothesis public layer plus all your private annotations, and annotations in private groups you belong to. If you don't authenticate you'll just search the public layer.

## Examples

### Find 10 recent notes, by `judell`, that have tags

```sql
select
  created,
  uri,
  tags,
  group_id,
  username
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
    username,
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
    id, username, created, text, uri, target
  order by
    id
)
select distinct
  'https://hypothes.is/a/' || id as link,
  *
from
  target_keys_to_rows
where
  target_key = 'Selector'
  and target->0->>'Selector' is null
order by
  created desc
```

### Find notes, in the Times' Opinion section, that quote selections matching "covid"

```sql
select
  'https://hypothes.is/a/' || id as link,
  uri,
  username,
  created,
  exact
from
  hypothesis_search
where
  query = 'wildcard_uri=https://www.nytimes.com/*/opinion/*'
  and exact ~* 'covid'
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
  r.username
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
    username,
    regexp_matches(uri, 'github.com/([^/]+)/([^/]+)') as match
  from
    hypothesis_search
  where
    query = 'wildcard_uri=http://github.com/*&limit=100'
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
    r.username
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
### Find URIs with conversational threads spanning more than one day

```
with thread_data as (
  select
    uri,
    count(*),
    min(created) as first,
    max(created) as last,
    sum(jsonb_array_length(refs)) as refs,
    array_agg(distinct username) as thread_participants
  from 
    hypothesis_search
  where 
    query = 'limit=1000' 
  group by uri	
)
select
  uri,
  count as annos,
  refs,
  first,
  last,
  date(last) - date(first) as days,
  thread_participants
from 
 thread_data
where
  date(last) - date(first) > 0
  and refs is not null
```

### Fetch the most recent 10000 annotations

```sql
select
  *
from
  hypothesis_search
where
  query = 'limit=10000'
```

**NOTE** When you use `limit` in the query string, it means: If there are `limit` annotations that match your query, fetch all of them. They will be stored in the Steampipe cache for 5 minutes by default, or longer if you add an `options` argument to your `hypothesis.spc` file and adjust the `cache_ttl` to a longer duration.

```
options "connection" {
  cache = "true"
  cache_ttl = 300 # default 5 minutes
}
```

#### Merging historical and live data

Suppose you have 500,000 annotations and are continuing to accumulate them at the rate of several thousand per day. (This is a real scenario.) You could stash the 500,000 in a table, or in a materialized view, like so:

```sql
create materialized view my_hypothesis_annotations as (
  select
    *
  from
    hypothesis_search
  where
    query = 'group=my_group&limit=500000'
) with data;
```

You could then merge those with live data like so.

```sql
with historical as (
  select
    *
  from
    my_hypothesis_annotations
),
new as (
  select
    *
  from
    hypothesis_search new
  where
    query = 'group=my_group&limit=5000'
    and not exists (
      select
        *
      from
        historical
      where
        historical.id = new.id
    )
),
select * from historical
union
select * from new
```
