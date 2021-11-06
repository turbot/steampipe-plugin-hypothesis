# Hypothesis Plugin for Steampipe

## Prerequisites

* [Steampipe](https://steampipe.io/downloads)

* [Golang](https://golang.org/doc/install)

## Build 

```sh
$ git clone https://github.com/judell/steampipe-plugin-hypothesis.git

$ cd steampipe-plugin-hypothesis

$ make # builds, then puts the plugin into your `~/.steampipe/plugins` directory

$ cp config/* ~/.steampipe/config # tells steampipe to load the plugin
```

## Try it!

```shell
$ steampipe query

> select 
    "user",
    tags
  from 
    hypothesis_search 
  where 
    query = 'uri=https://www.example.com'
  and jsonb_array_length(tags) > 0
  and "user" !~ 'judell'
]

   user   |                             tags
----------+--------------------------------------------------------------
 robins80 | ["rikersierra1"]
 robins80 | ["HypothesisTest", "3219099"]
 robins80 | ["HypothesisTest", "3219099"]
 ryany25  | ["asdf;", "asdfaasdf"]
 ryany25  | ["T-cell acute lymphoblastic leukemia-associated antigen 1"]
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

## More examples

[hypothesis_search](https://github.com/judell/steampipe-plugin-hypothesis/blob/main/docs/tables/hypothesis_search.md)

[hypothesis_profile](https://github.com/judell/steampipe-plugin-hypothesis/blob/main/docs/tables/hypothesis_profile.md)

## Links

Steampipe: [steampipe.io](https://steampipe.io)

Blog: [steampipe.io/blog](https://steampipe.io/blog)

Community: [steampipe.io/community/join](https://steampipe.io/community/join)

