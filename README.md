<p align="center">
  <h1 align="center">hypothesis Plugin for Steampipe</h1>
</p>

<p align="center">
  <a aria-label="Steampipe logo" href="https://steampipe.io">
    <img src="https://steampipe.io/images/steampipe_logo_wordmark_padding.svg" height="28">
  </a>
  <a aria-label="License" href="LICENSE">
    <img alt="" src="https://img.shields.io/static/v1?label=license&message=Apache-2.0&style=for-the-badge&labelColor=777777&color=F3F1F0">
  </a>
</p>

## Examples for creators of Steampipe plugins

Learn about [Steampipe](https://steampipe.io/)

## Get started

Install go, then:

```
$ git clone https://github.com/judell/steampipe-plugin-hypothesis

$ cp ./config/hypothesis.scp ~/.steampipe/config

$ make

$ steampipe query

> select * from hypothesis order by id
+----+----------+-------------------+
| id | greeting | json              |
+----+----------+-------------------+
| 1  | hypothesis    | {"hypothesis":"world"} |
| 2  | hypothesis    | {"hypothesis":"world"} |
| 3  | hypothesis    | {"hypothesis":"world"} |
+----+----------+-------------------+

```


