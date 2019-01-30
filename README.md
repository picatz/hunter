# `hunter`

Command-line application and golang client library for [hunter.io](https://hunter.io).

## Download

```console
$ go get -u github.com/picatz/hunter/...
```

## CLI Usage

The command-line application has three major commands `search`, `find`, and `verify`. All three of these commands output JSON. This makes parsing the infromation easy, especially using command-line tools like [`jq`](https://github.com/stedolan/jq).

```console
$ hunter
Usage:
  hunter [command]

Available Commands:
  find        Generates or retrieves the most likely email address from a domain name, a first name and a last name
  help        Help about any command
  search      Search all the email addresses corresponding to one website or company
  verify      Allows you to verify the deliverability of an email address

Flags:
  -h, --help   help for hunter

Use "hunter [command] --help" for more information about a command.
```

```console
$ hunter verify --email stevejobs@apple.com | jq -r .data.score
30
```

```console
$ hunter find --company google --full-name "Kelsey Hightower" | jq -r .data.email
khightower@google.com
```

```console
$ hunter search search --domain github.com --department it | jq -r '.data.emails[] | "\(.value) \(.position)"'
jbarnette@github.com Developer
bkeepers@github.com Developer
ben.balter@github.com Developer Product Manager
kyle.daigle@github.com Developer
vicent@github.com Principal Engineer Systems Engineering Manager
matthew@github.com Software Developer
kevin@github.com Developer
don@github.com Developer
john@github.com Developer
scott@github.com Developer
```

### `search`

```console
$ hunter search --help
...
$ hunter search --domain stripe.com
...
$ hunter search --company stripe
...
```

#### Output using `search`

```json
{
  "data": {
    "domain": "stripe.com",
    "disposable": false,
    "webmail": false,
    "pattern": "{first}",
    "organization": "Stripe",
    "emails": [
      ...
    ]
  },
  "meta": {
    ...
  }
}
```

### `find`

```console
$ hunter find --help
...
$ hunter find --domain asana.com --full-name "Dustin Moskovitz"
...
$ hunter find --company Asana --full-name "Dustin Moskovitz"
...
$ hunter find --company Asana --first-name Dustin --last-name Moskovitz
...
```

#### Output using `find`

```json
{
  "data": {
    "first_name": "Dustin",
    "last_name": "Moskovitz",
    "email": "dustin@asana.com",
    "score": 96,
    "domain": "asana.com",
    "position": "Cofounder",
    "twitter": "",
    "linkedin_url": "",
    "phone_number": "",
    "company": "Asana",
    "sources": [
      ...
    ]
  },
  "meta": {
    ...
  }
}
```

### `verify`

```console
$ hunter verify --help
...
$ hunter verify --email stevejobs@apple.com
...
```

#### Output using `verify`

```json
{
  "data": {
    "result": "undeliverable",
    "score": 30,
    "email": "stevejobs@apple.com",
    "regexp": true,
    "gibberish": false,
    "disposable": false,
    "webmail": false,
    "mx_records": true,
    "smtp_server": true,
    "smtp_check": false,
    "accept_all": false,
    "block": false,
    "sources": [
      ...
    ]
  },
  "meta": {
    ...
  }
}
```