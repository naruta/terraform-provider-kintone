# Terraform Provider for kintone

DISCLAIMER: This OSS is my own personal work and does not have any relationship with Cybozu Inc. or any other organization which I belong to.

## install

```
# MacOS
$ make install
```

## Usage

Create terraform file.
```hcl-terraform
provider "kintone" {
  host = <KINTONE_HOST>
  user = <KINTONE_USER>
  password = <KINTONE_PASSWORD>
}
resource "kintone_application" {
  name = "Terraform MyApp"
  description = "This is my application!!"
  theme = "BLACK"

  field = { code = "code_number" label = "number" type = "NUMBER" }
  field = { code = "code_title" label = "title" type = "SINGLE_LINE_TEXT" }
}
```

Discovery Plugin
```
$ terraform init
```

Check Plan
```
$ terraform plan
```

## License
MIT license.
