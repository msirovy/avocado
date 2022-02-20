variable "loki_url" {
    type = string
    default = ""
}

job "avocado" {
    datacenters = ["home"]

    group "garden" {
        task "avocado" {
            driver = "raw_exec"

            config {
                command = "/usr/local/bin/avocado"
                args = [""]
            }

            env {
                LOKI_URL = var.loki_url
            }

            /*
            artifact {
                source = "https://internal.file.server/name-of-my-binary"
                options {
                    checksum = "sha256:abd123445ds4555555555"
                }
            }
            */

        }
    }
}