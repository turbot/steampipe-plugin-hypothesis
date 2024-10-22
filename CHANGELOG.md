## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#46](https://github.com/turbot/steampipe-plugin-hypothesis/pull/46))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#46](https://github.com/turbot/steampipe-plugin-hypothesis/pull/46))

## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#33](https://github.com/turbot/steampipe-plugin-hypothesis/pull/33))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#33](https://github.com/turbot/steampipe-plugin-hypothesis/pull/33))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hypothesis/blob/main/docs/LICENSE). ([#33](https://github.com/turbot/steampipe-plugin-hypothesis/pull/33))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#32](https://github.com/turbot/steampipe-plugin-hypothesis/pull/32))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#25](https://github.com/turbot/steampipe-plugin-hypothesis/pull/25))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#22](https://github.com/turbot/steampipe-plugin-hypothesis/pull/22))
- Recompiled plugin with Go version `1.21`. ([#22](https://github.com/turbot/steampipe-plugin-hypothesis/pull/22))

## v0.4.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#17](https://github.com/turbot/steampipe-plugin-hypothesis/pull/17))

## v0.3.1 [2022-10-03]

_Dependencies_

- Recompiled plugin with [hypothesis-go v0.2.6](https://github.com/judell/hypothesis-go/releases/tag/v0.2.6) which includes a regression fix.

## v0.3.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#14](https://github.com/turbot/steampipe-plugin-hypothesis/pull/14))

## v0.2.0 [2022-08-29]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v414-2022-08-26) which includes several caching and memory management improvements. ([#11](https://github.com/turbot/steampipe-plugin-hypothesis/pull/11))
- Recompiled plugin with Go version `1.19`. ([#11](https://github.com/turbot/steampipe-plugin-hypothesis/pull/11))

## v0.1.1 [2022-08-18]

- Recompile plugin with `github.com/judell/hypothesis-go@v0.2.4` which handles a breaking change in the Hypothesis API. ([#10](https://github.com/turbot/steampipe-plugin-hypothesis/pull/10))

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
