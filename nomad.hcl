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
                command = "avocado"
                args = ["-interval", "300"]
            }

            env {
                LOKI_URL = var.loki_url
            }

            
            artifact {
                source = "https://github.com/msirovy/avocado/releases/download/0.1.2/avocado_0.1.2_linux_armv7.tar.gz"
                /*
                options {
                    checksum = "sha256:abd123445ds4555555555"
                }
                */
            }
            

        }
    }
}