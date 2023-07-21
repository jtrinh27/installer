locals {

  // Because of the issue https://github.com/hashicorp/terraform/issues/12570, the consumers cannot count 0/1
  // based on if api_external_lb_dns_name for example, which will be null when there is no external lb for API.
  // So publish_strategy serves an coordinated proxy for that decision.
  public_endpoints = var.publish_strategy == "External" ? true : false

  use_cname = contains(["us-gov-west-1", "us-gov-east-1"], var.region)
  use_alias = ! local.use_cname
}

provider "aws" {
  alias = "private_hosted_zone"

  assume_role {
    role_arn = var.internal_zone_role
  }

  region = var.region

  skip_region_validation = true

  endpoints {
    ec2     = lookup(var.custom_endpoints, "ec2", null)
    elb     = lookup(var.custom_endpoints, "elasticloadbalancing", null)
    iam     = lookup(var.custom_endpoints, "iam", null)
    s3      = lookup(var.custom_endpoints, "s3", null)
    sts     = lookup(var.custom_endpoints, "sts", null)
  }
}