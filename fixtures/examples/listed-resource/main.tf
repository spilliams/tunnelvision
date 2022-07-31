variable "qty" {
  type    = number
  default = 0
}

resource "random_string" "this" {
  count = var.qty

  length = 16

  upper            = false
  lower            = false
  numeric          = true
  special          = true
  override_special = "abcdef"
}

output "hexes" {
  value = [for s in random_string.this : s.result]
}
