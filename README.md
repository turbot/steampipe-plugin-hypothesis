![image](https://hub.steampipe.io/images/plugins/turbot/hypothesis-social-graphic.png)
# Hypothesis Plugin for Steampipe

Use SQL to query Hypothesis annotations.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/hypothesis)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/hypothesis/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-hypothesis)
## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install hypothesis
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/hypothesis#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/hypothesis#configuration).

Run a query:

```sql
  select 
    "user",
    created,
    tags
  from 
    hypothesis_search 
  where 
    query = 'uri=https://www.example.com';
```
## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-hypothesis.git
cd steampipe-plugin-hypothesis
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
emacs ~/.steampipe/config/hypothesis.spc
```

Try it!

```shell
steampipe query
> .inspect hypothesis
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-twilio/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Twilio Plugin](https://github.com/turbot/steampipe-plugin-hypothesis/labels/help%20wanted)



