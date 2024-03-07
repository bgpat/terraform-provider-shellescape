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
