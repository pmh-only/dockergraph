# :whale2: `dockergraph`
Renders fancy(?) docker network graph powered by Go & Graphviz

![](/docs/graph.svg)
> :link: demo: https://shutupandtakemy.codes/#netgraph

## features.
* Draws `container` and `network_gateway` nodes.
* Provides `sha256 id`, `name` and `image url` informations.
* Highlights containers with `db`, `redis`, `nginx` in its name.
* Supports `.svg` and `.png` rendering!

## `docker-compose.yml`
```yml
version: '3'

services:
  netgraph:
    image: shutupandtakemy.codes/dockergraph
    network_mode: none
    restart: unless-stopped
    user: root
    command: sh -c 'while true; do echo "Running..."; ./dockergraph; sleep 30; done'
    environment:
      - DOCKER_API_VERSION=1.41 # change me
    volumes:
      - /var/www/html/graph.svg:/app/graph.svg:rw
      - /var/www/html/graph.png:/app/graph.png:rw
      - /var/run/docker.sock:/var/run/docker.sock:ro
```

## contribution.
* MIT Licensed.
* feel free to send me PR!
