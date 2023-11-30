---
title: "Steampipe Table: hypothesis_search - Query Hypothesis Searches using SQL"
description: "Allows users to query Hypothesis Searches, specifically the annotations and their corresponding details, providing insights into user annotations and potential patterns."
---

# Table: hypothesis_search - Query Hypothesis Searches using SQL

Hypothesis is a service that allows users to annotate web pages and PDFs, fostering conversations within the text. It is used by educators, journalists, publishers, and researchers to anchor discussions, express opinions, and share insights directly on top of digital content. Hypothesis Searches are a resource within this service that allows users to query and retrieve these annotations.

## Table Usage Guide

The `hypothesis_search` table provides insights into Hypothesis Searches within the Hypothesis service. As a researcher or educator, explore annotation-specific details through this table, including the text, tags, and user who made the annotation. Utilize it to uncover information about annotations, such as those with specific tags, the users who made them, and their corresponding details.

## Examples

### Find 10 recent notes, by `judell`, that have tags
Explore the recent activity of a specific user, 'judell', to identify instances where they have added tags to their notes. This is useful for understanding their areas of interest and their tagging habits.

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
Explore the instances where notes are tagged with both 'media' and 'review', allowing you to focus on specific areas of interest. This can be particularly useful when you're looking for overlaps in topics or themes.
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
Explore the instances where notes are tagged with both 'social media' and 'peer review', which can be useful for understanding the intersection of these two topics in your data.

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
Explore the frequency of annotations made on the New York Times homepage in a given month. This can help you understand the level of engagement or significant events during specific periods.

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
Explore which articles in the Times' Opinion section have been annotated and understand the frequency of these annotations. This can be useful to identify the most discussed or controversial articles.

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
Determine the instances where notes were made on the entire webpage of www.example.com, rather than a specific selection. This is useful for identifying overall feedback or comments about the webpage as a whole.

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
Discover the segments that quote selections matching a specific term within the Times' Opinion section. This is useful for exploring user-generated annotations and comments on current events or trending topics, such as 'covid'.

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
Explore annotated GitHub repositories and gain insights into their associated details from the GitHub API. This can help you identify patterns or trends in the data, enhancing your understanding of these repositories.
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
This query is used to analyze the frequency of annotations made by users on different GitHub repositories. It can be beneficial to identify which repositories are receiving the most attention or interaction, potentially indicating areas of high interest or activity.

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
Determine the areas in which the most recent 10,000 annotations exist. This is useful for gaining insights into the latest trends and patterns in your data, especially when dealing with large volumes of annotations that are continuously accumulating.

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