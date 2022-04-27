## v0.1.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#6](https://github.com/turbot/steampipe-plugin-hypothesis/pull/6))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#5](https://github.com/turbot/steampipe-plugin-hypothesis/pull/5))

## v0.0.2 [2022-02-28]

_What's new?_

- New column added: `refs`. Enables conversational analysis. See [example](https://hub.steampipe.io/plugins/turbot/hypothesis/tables/hypothesis_search#find-uris-with-conversational-threads-spanning-more-than-one-day).


## v0.0.1 [2022-01-14]

_What's new?_

- New tables added

  - [hypothesis_profile](https://hub.steampipe.io/plugins/turbot/hypothesis/tables/hypothesis_profile)
  - [hypothesis_search](https://hub.steampipe.io/plugins/turbot/hypothesis/tables/hypothesis_search)
