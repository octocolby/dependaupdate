terraform {
  required_version = ">= 1.5.0"

  required_providers {
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

resource "null_resource" "demo" {
  provisioner "local-exec" {
    command = "echo 'Terraform demo is working!'"
  }
}

output "demo_message" {
  value = "Terraform apply completed successfully"
}
