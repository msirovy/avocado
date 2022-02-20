# AVOCADO






Deploy avocado to nomad this way

```bash
nomad run -var 'loki_url=https://__ID__:__PASSWORD__=@logs-prod-us-central1.grafana.net/api/prom/push' avocado.hcl

```