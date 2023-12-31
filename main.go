package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func main() {
	g := graphviz.New()
	template_graph := []byte("digraph{\noverlap_scaling=-7\n}")

	graph, _ := graphviz.ParseBytes(template_graph)
	cli, _ := client.NewClientWithOpts(client.FromEnv)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	networks := map[string]*cgraph.Node{}
	
	for _, container := range containers {
		container_name := fmt.Sprintf("%s\n%s\n%s", container.ID[:8], container.Names[0], container.Image)
		container_node, _ := graph.CreateNode(container_name)
		container_node.SetShape(cgraph.CircleShape)
		container_node.SetHeight(2)
		container_node.SetWidth(2)
		container_node.SetStyle(cgraph.FilledNodeStyle)
		container_node.SetFillColor("#FAA4C3")
		container_node.SetTooltip(container_name)

		if strings.Contains(container.Names[0], "db") {
			container_node.SetShape(cgraph.CylinderShape)
			container_node.SetFillColor("#FAEA99")
		}

		if strings.Contains(container.Names[0], "redis") {
			container_node.SetShape(cgraph.CylinderShape)
			container_node.SetFillColor("#FA707D")
		}

		if strings.Contains(container.Names[0], "nginx") {
			container_node.SetShape(cgraph.Box3DShape)
			container_node.SetFillColor("#9EFAAF")
		}

		for network, networkDetail := range container.NetworkSettings.Networks {
			if _, ok := networks[network]; !ok {
				network_label := fmt.Sprintf("%s\n%s\n%s", networkDetail.NetworkID[:8], network, networkDetail.Gateway)
				network_node, _ := graph.CreateNode(network_label)
				network_node.SetStyle(cgraph.FilledNodeStyle)
				network_node.SetShape(cgraph.BoxShape)
				network_node.SetFillColor("#9EA1FA")
				networks[network] = network_node
			}

			// ignore maxscale connection
			if network == "maxscale" {
				continue
			}

			edge, _ := graph.CreateEdge("", container_node, networks[network])
			edge.SetTailLabel(networkDetail.IPAddress)
		}
	}

	g.SetLayout(graphviz.SFDP)

	err = g.RenderFilename(graph, graphviz.SVG, "graph.svg")
	if err != nil {
		panic(err)
	}

	err = g.RenderFilename(graph, graphviz.PNG, "graph.png")
	if err != nil {
		panic(err)
	}
}
