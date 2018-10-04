package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	uuid "github.com/satori/go.uuid"
)

const (
	oneToOne   = "1..1"
	oneToMany  = "1..n"
	manyToMany = "n..n"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// GraphDefinition holds the initial and final graph structure
type GraphDefinition struct {
	NodeDefs            []NodeDefinition         `json:"nodes"`
	RelationshipDefs    []RelationshipDefinition `json:"relationships"`
	RemoveOrphanedNodes bool                     `json:"removeOrphanedNodes"`
}

// PropertiesDefinition holds the available properties on either a node or an edge
type PropertiesDefinition map[string]Property

// NodeDefinition holds a node object
type NodeDefinition struct {
	Count      int                  `json:"count"`
	Properties PropertiesDefinition `json:"properties"`
	Label      string               `json:"label"`
}

// RelationshipDefinition defines an edge between nodes
type RelationshipDefinition struct {
	Source     string               `json:"source"`
	Target     string               `json:"target"`
	Label      string               `json:"label"`
	Properties PropertiesDefinition `json:"properties"`
	Min        int                  `json:"min"`
	Max        int                  `json:"max"`
	Chance     int                  `json:"probability"`
	Type       string               `json:"type"`
}

// Node ...
type Node struct {
	Properties map[string]interface{} `json:"properties"`
	Label      string                 `json:"label"`
	ID         string                 `json:"id"`
}

// Edge ...
type Edge struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	Label       string `json:"label"`
	SourceLabel string `json:"sourceLabel"`
	TargetLabel string `json:"targetLabel"`
}

// Graph ...
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func generateProperties(p PropertiesDefinition) map[string]interface{} {
	props := map[string]interface{}{}
	for propertyKey, property := range p {
		props[propertyKey] = property.Populate()
	}
	return props
}

func verifyRelationshipDefinition(rd RelationshipDefinition, nodesByLabel map[string][]Node) error {
	if rd.Source == "" {
		return errors.New("relationship definition must include source")
	}
	if rd.Target == "" {
		return errors.New("relationship definition must include target")
	}
	if nodesByLabel[rd.Source] == nil {
		return errors.New("missing node definition for label: " + rd.Source)
	}
	if nodesByLabel[rd.Target] == nil {
		return errors.New("missing node definition for label: " + rd.Target)
	}
	if rd.Type != oneToMany && rd.Type != manyToMany && rd.Type != oneToOne {
		return errors.New("relationship definition must include type: 1..n, 1..1, n..n")
	}
	if rd.Label == "" {
		return errors.New("relationship definition must include label")
	}
	return nil
}

func shuffleNodes(vals []Node) []Node {
	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
	return vals
}

func createNodes(g GraphDefinition) (map[string][]Node, error) {
	nodesByLabel := map[string][]Node{}
	for _, nodeDef := range g.NodeDefs {
		nodesByLabel[nodeDef.Label] = []Node{}
		for i := 0; i < nodeDef.Count; i++ {
			id, err := uuid.NewV4()
			if err != nil {
				return nodesByLabel, errors.New("unable to generate uuid for node")
			}
			generatedNode := Node{
				Label:      nodeDef.Label,
				Properties: generateProperties(nodeDef.Properties),
				ID:         id.String(),
			}
			nodesByLabel[nodeDef.Label] = append(nodesByLabel[nodeDef.Label], generatedNode)
		}
	}
	return nodesByLabel, nil
}

func addEdge(edges []Edge, u, v Node, label string) []Edge {
	return append(edges, Edge{
		Source:      u.ID,
		Target:      v.ID,
		Label:       label,
		SourceLabel: u.Label,
		TargetLabel: v.Label,
	})
}

func buildGraph(g GraphDefinition) (Graph, error) {
	graph := Graph{}

	nodesByLabel, err := createNodes(g)
	if err != nil {
		return graph, err
	}
	edges := []Edge{}

	// No Set types so we can use a map and loop over the values at the end
	// to only get the nodes that have edges
	nodes := map[string]Node{}
	for _, d := range g.RelationshipDefs {
		if err := verifyRelationshipDefinition(d, nodesByLabel); err != nil {
			return graph, err
		}
		sources := nodesByLabel[d.Source]
		targetConnects := map[string]int{}
		sourceConnects := map[string]int{}

		for _, u := range sources {
			targets := shuffleNodes(nodesByLabel[d.Target])
			for _, v := range targets {
				if (d.Chance != 0 && random(0, 100) < d.Chance) || d.Chance == 0 {
					if (d.Max > 0 && sourceConnects[u.ID] < d.Max) || d.Max == 0 {
						if d.Type == oneToMany {
							if targetConnects[v.ID] == 0 {
								sourceConnects[u.ID]++
								targetConnects[v.ID]++
								edges = addEdge(edges, u, v, d.Label)
								nodes[u.ID] = u
								nodes[v.ID] = v
							}
							continue
						}
						if d.Type == oneToOne {
							if sourceConnects[u.ID] != 0 {
								break
							}
							if targetConnects[v.ID] != 0 {
								continue
							}
							sourceConnects[u.ID]++
							targetConnects[v.ID]++
							edges = addEdge(edges, u, v, d.Label)
							nodes[u.ID] = u
							nodes[v.ID] = v
						}
						if d.Type == manyToMany {
							sourceConnects[u.ID]++
							targetConnects[v.ID]++
							edges = addEdge(edges, u, v, d.Label)
							nodes[u.ID] = u
							nodes[v.ID] = v
						}
					}
				}
			}

			// We want to remove targets that have not been connected
			// this will make sure that we get a connected graph
			connectedTargets := []Node{}
			for _, node := range targets {
				if targetConnects[node.ID] > 0 {
					connectedTargets = append(connectedTargets, node)
				}
			}
			nodesByLabel[d.Target] = connectedTargets
		}
	}

	nodeArray := []Node{}
	for _, node := range nodes {
		nodeArray = append(nodeArray, node)
	}
	graph.Nodes = nodeArray
	graph.Edges = edges
	return graph, nil
}

type lambdaError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var headers = map[string]string{
	"Access-Control-Allow-Origin": "*",
}

func handleError(errorObject error, code int) (events.APIGatewayProxyResponse, error) {
	e := events.APIGatewayProxyResponse{}
	l := lambdaError{
		Message: errorObject.Error(),
		Code:    code,
	}

	b, err := json.Marshal(l)
	if err != nil {
		return e, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(b),
		Headers:    headers,
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	g := GraphDefinition{}
	if err := json.Unmarshal([]byte(request.Body), &g); err != nil {
		return handleError(err, 400)
	}

	graph, err := buildGraph(g)
	if err != nil {
		return handleError(err, 400)
	}

	mGraph, err := json.Marshal(graph)
	if err != nil {
		return handleError(err, 400)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(mGraph),
		StatusCode: 200,
		Headers:    headers,
	}, nil
}

func local() {
	jsonFile, err := os.Open("schema.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	graph := GraphDefinition{}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

	if err := json.Unmarshal(byteValue, &graph); err != nil {
		log.Fatal(err)
	}
	g, err := buildGraph(graph)
	bytes, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))
}

func main() {
	if _, exists := os.LookupEnv("IS_LAMBDA"); exists {
		lambda.Start(handler)
	} else {
		local()
	}
}
