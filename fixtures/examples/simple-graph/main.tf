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

locals {
  full_name = "tf-${var.name}-${random_string.slug.result}"
}

output "full_name" {
  value = local.full_name
}
