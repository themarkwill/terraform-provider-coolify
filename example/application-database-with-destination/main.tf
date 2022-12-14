terraform {
  required_providers {
    coolify = {
      source  = "themarkwill/coolify"
      version = "0.4.4"
    }
  }
}

provider "coolify" {
  address = "url of coolify"
  token = "token"
}

resource "coolify_destination" "network" {
  name    = "Seccond Application Terraform"
  network = "second-network-bibi"
}

resource "coolify_database" "redis" {
  name   = "second-application-db"
  engine = "redis:7.0"

  settings {
    destination_id = coolify_destination.network.id
    is_public      = true
    password       = "123456"
  }
}

resource "coolify_application" "test_item" {
  name   = "second-app"
  domain = "https://second-app.s.coolify.io"

  template {
    build_pack  = "node"
    image       = "node:14"
    build_image = "node:14"

    settings {
      install_command = "npm install"
      start_command   = "npm start"
      auto_deploy     = false
    }

    env {
      key   = "BASE_PROJECT"
      value = "production"
    }

    env {
      key   = "REDIS_PASSWORD"
      value = coolify_database.redis.status.password
    }

    env {
      key   = "REDIS_HOST"
      value = coolify_database.redis.status.host
    }

    env {
      key   = "REDIS_PORT"
      value = coolify_database.redis.status.port
    }
  }

  repository {
    repository_id = 579493141
    repository    = "cool-sample/sample-nodejs"
    branch        = "develop"
  }

  settings {
    destination_id = coolify_destination.network.id
    source_id      = "clb9y09gs000f9dmod69f7dce"
  }
}