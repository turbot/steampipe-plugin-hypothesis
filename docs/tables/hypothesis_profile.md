---
title: "Steampipe Table: hypothesis_profile - Query Hypothesis Profiles using SQL"
description: "Allows users to query Hypothesis Profiles, specifically data related to profile id, user id, and preferences, providing insights into user profiles and preferences."
---

# Table: hypothesis_profile - Query Hypothesis Profiles using SQL

Hypothesis is an open-source, community-driven platform that facilitates annotation and discussion on web content. Profiles in Hypothesis represent individual user accounts, storing data related to user id, preferences, and other related information. Through these profiles, Hypothesis provides a platform for users to engage in collaborative annotation and discussion.

## Table Usage Guide

The `hypothesis_profile` table provides insights into user profiles within the Hypothesis platform. As a community manager or moderator, explore user-specific details through this table, including user id and preferences. Utilize it to uncover information about users, such as their annotation preferences and activity, facilitating more effective community management and engagement.

## Examples

### Get your username, display name, and authority
Explore your Hypothesis profile to uncover your username, display name, and authority level. This can help in understanding your user status and permissions within the platform.

```sql
select
  username,
  display_name,
  authority
from
  hypothesis_profile
```

### Get the names and ids of groups you belong to
Explore which groups you are a part of, in order to better understand your affiliations and interactions within the platform. This can be useful in managing your group memberships and identifying areas for collaboration.

```sql
select
  jsonb_array_elements(groups)
from
  hypothesis_profile
```

### Among the most recent 500 notes, find those in your private groups (method 1)
Determine the areas in which your private group notes are among the most recent 500. This is useful for prioritizing review and response to the most recent discussions within your private groups.

```sql
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
    query = 'limit=500'
  order by
    created desc
)
select
  'https://hypothes.is/a/' || a.id as link,
  g.group_info ->> 'name' as name,
  a.username,
  a.created,
  a.title,
  a.uri
from
  groups g
join
  annos a
on
  g.group_info ->> 'id' = a.group_id
where
  g.group_info ->> 'public' != 'true'
```

### Among the most recent 500 notes, find those in your private groups (method 2)
Determine the instances of your most recent private group notes within the last 500 entries. This is useful for reviewing and managing your private group content without having to sift through all notes manually.
**NOTE** It can be helpful to turn chunks of SQL code into Postgres functions. Here we define, and then use, `hypothesis_is_private_group`, a function that checks if a `group_id` is private. This function makes method 2 simpler than method 1. And you can use the function anywhere a `group_id` appears. See [Postgres functional style](https://blog.jonudell.net/2021/08/21/postgres-functional-style/) for details.

#### Create the function `hypothesis_is_private_group`

**NOTE** Steampipe plugins put their tables into Postgres schemas (namespaces) that match the names of the plugins. So tables in this plugin are actually `hypothesis.hypothesis_search` and `hypothesis.hypothesis_profile`. The examples here don't qualify table names with schemas because if there is no confict with another schema it's unnecessary. When you create functions, though, they live in Postgres' global namespace. Nothing requires you to prepend a schema-like prefix to function names, but it's probably a good idea to do that in order to clarify which plugins they're intended to work with.


```sql
create function hypothesis_is_private_group (group_id text) returns boolean as $$
  declare is_private boolean;
  begin
    with groups as (
    select
        jsonb_array_elements(groups) as group_info
    from
        hypothesis_profile
    )
    select
      g.group_info ->> 'public' != 'true'
    from
      groups g
    where
      g.group_info ->> 'id' = group_id
    into
      is_private;
    return is_private;
  end;
$$ language plpgsql;
```

#### Use `hypothesis_is_private_group`

```sql
select
  'https://hypothes.is/a/' || id as link,
  group_id,
  username,
  created,
  title,
  uri
from
  hypothesis_search
where
  query = 'limit=500'
  and hypothesis_is_private_group(group_id)
```