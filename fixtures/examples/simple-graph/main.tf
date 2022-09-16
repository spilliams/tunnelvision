variable "name" {
  type    = string
  default = "myResource"
}

resource "random_string" "slug" {
  length = 4

  upper            = false
  lower            = false
  numeric          = true
  special          = true
  override_special = "abcdef"
}

module "hex" {
  source = "../../modules/hexadecimal"
}

data "local_file" "this" {
  filename = "${path.module}/main.tf"
}

locals {
  full_name = "tf-${var.name}-${random_string.slug.result}"
}

output "full_name" {
  value = local.full_name
}

output "file_length" {
  value = length(data.local_file.this.content)
}
