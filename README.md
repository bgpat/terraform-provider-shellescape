# Terraform Shell Escape Data Source

![release](https://github.com/bgpat/terraform-provider-shellescape/workflows/release/badge.svg)

A Terraform data source provider to escape shell special characters.

## Usage

```terraform
data "shellescape_quote" "hello" {
  string = <<-EOT
  _______
  < hello >
  -------
          \   ^__^
           \  (oo)\_______
              (__)\       )\/\
                  ||----w |
                  ||     ||
  EOT
}

resource "local_file" "hello" {
    content     = "echo ${data.shellescape_quote.hello.quoted}"
    filename = "hello.sh"
}
```

## Installation

1. Build from source

   ```bash
   go get -d github.com/bgpat/terraform-provider-shellescape
   ```

2. Place into plugins directory

   ```bash
   mkdir -p ~/.terraform.d/plugins
   mv $GOPATH/bin/terraform-provider-shellescape ~/.terraform.d/plugins/terraform-provider-shellescape
   ```
