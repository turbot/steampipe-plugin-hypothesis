---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/hypothesis.svg"
brand_color: "#BD1C2B"
display_name: "Hypothesis"
short_name: "hypothesis"
description: "Steampipe plugin to query Hypothesis annotations."
og_description: "Query Hypothesis with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/hypothesis-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---
# Hypothesis + Steampipe

[Hypothesis](https://hypothes.is) is a web annnotation system.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List annotations on `www.example.com`, with at least one tag, by a user other than `judell`:

```sql
select
  username,
  tags
from
  hypothesis_search
where
  query = 'uri=https://www.example.com'
and jsonb_array_length(tags) > 0
and username !~ 'judell'
```

```shell
+----------+--------------------------------------------------------------+
| username | tags                                                         |
+----------+--------------------------------------------------------------+
| ryany25  | ["asdf;","asdfaasdf"]                                        |
| ryany25  | ["T-cell acute lymphoblastic leukemia-associated antigen 1"] |
| robins80 | ["rikersierra1"]                                             |
| robins80 | ["HypothesisTest","3219099"]                                 |
| robins80 | ["HypothesisTest","3219099"]                                 |
+----------+--------------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/hypothesis/tables)**

## Get started

### Install

Download and install the latest Hypothesis plugin:

```bash
steampipe plugin install hypothesis
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Get your API token from the [Hypothesis service](https://hypothes.is/account/developer) (you must be logged in to Hypothesis to view this page). The token is optional. Without it, you can still query the Hypothesis public layer.

### Configuration

Installing the latest hypothesis plugin will create a config file (`~/.steampipe/config/hypothesis.spc`) with a single connection named `hypothesis`:

If you are a Hypothesis user wanting to query your own private notes, or notes in private groups you belong to, then uncomment `#token` and provide your API token.

```hcl
connection "hypothesis" {
  plugin = "hypothesis"
  #token = "6879-35....3df5"
}
```


