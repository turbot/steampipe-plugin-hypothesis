# Hypothesis

The Hypothesis plugin queries annotations stored on the [Hypothesis](https://hypothes.is) service.

## Installation

```bash
$ git clone https://github.com/judell/steampipe-plugin-hypothesis.git

$ cd hypothesis-go

$ make

$ cp config/* ~/.steampipe/config
```

## API token

The token is optional. Without it, you can still query the Hypothesis public layer. 

If you are a Hypothesis user wanting to query your own private notes, or notes in private groups you belong to, then log in, open https://hypothes.is/account/developer, generate a token, and copy it into `~/.steampipe/config/hypothesis.spc` like so.

  ```hcl
  connection "hypothesis" {
    plugin  = "hypothesis"
    token   = "6879-35....3df5"
  }
```
