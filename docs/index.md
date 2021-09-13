---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d1a"
display_name: Hello
name: hello
description: Steampipe plugin to show how to write plugins
---

# Hello

[Hello]() is a series of examples showing how to write Steampipe plugins


## Installation

Build and use locally.


## Tables

```
> .inspect hello
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| hello_1            | Simplest possible way to populate a query with data                         |
+--------------------+-----------------------------------------------------------------------------+
```

## Credentials

None.

## Connection Configuration

Put this into ~/.steampipe/config/hello.spc`.

```
connection "hello" {
  plugin    = "hello"
}
```
