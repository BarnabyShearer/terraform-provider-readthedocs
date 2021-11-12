terraform {
  required_providers {
    readthedocs = {
      source = "BarnabyShearer/readthedocs"
    }
  }
}

provider "readthedocs" {
  token = "TOKEN" # optionally use READTHEDOCS_TOKEN env var
}
