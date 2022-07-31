resource "random_string" "this" {
  length = 16

  upper            = false
  lower            = false
  numeric          = true
  special          = true
  override_special = "abcdef"
}

output "hex" {
  value = random_string.this.result
}
