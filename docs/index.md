---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d1a"
display_name: hypothesis
name: hypothesis
description: Steampipe plugin to show how to write plugins
---

# hypothesis

[hypothesis]() is a series of examples showing how to write Steampipe plugins


## Installation

Build and use locally.


## Tables

```
> .inspect hypothesis
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| hypothesis              | Simplest possible way to populate a query with data                         |
+--------------------+-----------------------------------------------------------------------------+
```

## Credentials

None.

## Connection Configuration

Put this into ~/.steampipe/config/hypothesis.spc`.

```
connection "hypothesis" {
  plugin    = "hypothesis"
}
```
