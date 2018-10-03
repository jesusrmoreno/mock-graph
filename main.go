package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/icrowley/fake"
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

func populate(prop string) interface{} {
	p := strings.Title(prop)
	switch p {
	case "Name":
		v := []string{fake.MaleFullName(), fake.FemaleFullName()}
		return v[random(0, 2)]
	case "Brand":
		return fake.Brand()
	case "Character":
		return fake.Character()
	case "Characters":
		return fake.Characters()
	case "City":
		return fake.City()
	case "Color":
		return fake.Color()
	case "Company":
		return fake.Company()
	case "Continent":
		return fake.Continent()
	case "Country":
		return fake.Country()
	case "CreditCardNum":
		return fmt.Sprintf("0000-%d-%d-%d", random(1000, 10000), random(1000, 10000), random(1000, 10000))
	case "CreditCardType":
		return fake.CreditCardType()
	case "Currency":
		return fake.Currency()
	case "CurrencyCode":
		return fake.CurrencyCode()
	case "Day":
		return fake.Day()
	case "Digits":
		return fake.Digits()
	case "DomainName":
		return fake.DomainName()
	case "DomainZone":
		return fake.DomainZone()
	case "EmailAddress":
		return fake.EmailAddress()
	case "EmailBody":
		return fake.EmailBody()
	case "EmailSubject":
		return fake.EmailSubject()
	case "FemaleFirstName":
		return fake.FemaleFirstName()
	case "FemaleFullName":
		return fake.FemaleFullName()
	case "FemaleFullNameWithPrefix":
		return fake.FemaleFullNameWithPrefix()
	case "FemaleFullNameWithSuffix":
		return fake.FemaleFullNameWithSuffix()
	case "FemaleLastName":
		return fake.FemaleLastName()
	case "FemalePatronymic":
		return fake.FemalePatronymic()
	case "FirstName":
		return fake.FirstName()
	case "FullName":
		return fake.FullName()
	case "FullNameWithPrefix":
		return fake.FullNameWithPrefix()
	case "FullNameWithSuffix":
		return fake.FullNameWithSuffix()
	case "Gender":
		return fake.Gender()
	case "GenderAbbrev":
		return fake.GenderAbbrev()
	case "GetLangs":
		return fake.GetLangs()
	case "HexColor":
		return fake.HexColor()
	case "HexColorShort":
		return fake.HexColorShort()
	case "IPv4":
		return fake.IPv4()
	case "IPv6":
		return fake.IPv6()
	case "Industry":
		return fake.Industry()
	case "JobTitle":
		return fake.JobTitle()
	case "Language":
		return fake.Language()
	case "LastName":
		return fake.LastName()
	case "Latitude":
		return fake.Latitude()
	case "LatitudeDegrees":
		return fake.LatitudeDegrees()
	case "LatitudeDirection":
		return fake.LatitudeDirection()
	case "LatitudeMinutes":
		return fake.LatitudeMinutes()
	case "LatitudeSeconds":
		return fake.LatitudeSeconds()
	case "Longitude":
		return fake.Longitude()
	case "LongitudeDegrees":
		return fake.LongitudeDegrees()
	case "LongitudeDirection":
		return fake.LongitudeDirection()
	case "LongitudeMinutes":
		return fake.LongitudeMinutes()
	case "LongitudeSeconds":
		return fake.LongitudeSeconds()
	case "MaleFirstName":
		return fake.MaleFirstName()
	case "MaleFullName":
		return fake.MaleFullName()
	case "MaleFullNameWithPrefix":
		return fake.MaleFullNameWithPrefix()
	case "MaleFullNameWithSuffix":
		return fake.MaleFullNameWithSuffix()
	case "MaleLastName":
		return fake.MaleLastName()
	case "MalePatronymic":
		return fake.MalePatronymic()
	case "Model":
		return fake.Model()
	case "Month":
		return fake.Month()
	case "MonthNum":
		return fake.MonthNum()
	case "MonthShort":
		return fake.MonthShort()
	case "Paragraph":
		return fake.Paragraph()
	case "Paragraphs":
		return fake.Paragraphs()
	case "Patronymic":
		return fake.Patronymic()
	case "Phone":
		return fake.Phone()
	case "Product":
		return fake.Product()
	case "ProductName":
		return fake.ProductName()
	case "Sentence":
		return fake.Sentence()
	case "Sentences":
		return fake.Sentences()
	case "SimplePassword":
		return fake.SimplePassword()
	case "State":
		return fake.State()
	case "StateAbbrev":
		return fake.StateAbbrev()
	case "Street":
		return fake.Street()
	case "StreetAddress":
		return fake.StreetAddress()
	case "Title":
		return fake.Title()
	case "TopLevelDomain":
		return fake.TopLevelDomain()
	case "UserAgent":
		return fake.UserAgent()
	case "UserName":
		return fake.UserName()
	case "WeekDay":
		return fake.WeekDay()
	case "WeekDayShort":
		return fake.WeekDayShort()
	case "WeekdayNum":
		return fake.WeekdayNum()
	case "Word":
		return fake.Word()
	case "Words":
		return fake.Words()
	case "Year":
		return fake.Year(0, time.Now().Year())
	case "Zip":
		return fake.Zip()
	}
	return prop
}

// GraphDefinition holds the initial and final graph structure
type GraphDefinition struct {
	NodeDefs            []NodeDefinition         `json:"nodes"`
	RelationshipDefs    []RelationshipDefinition `json:"relationships"`
	RemoveOrphanedNodes bool                     `json:"removeOrphanedNodes"`
	Entries             []EntryDefinition        `json:"entries"`
}

// EntryDefinition ...
type EntryDefinition struct {
	Label               string `json:"label"`
	IncludeDisconnected string `json:"includeDisconnected"`
}

// PropertiesDefinition holds the available properties on either a node or an edge
type PropertiesDefinition map[string]string

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
	for prop, kind := range p {
		props[prop] = populate(kind)
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

func connectedNodes(nodes []Node, connectedLookup map[string]bool) []Node {
	connected := []Node{}
	for _, node := range nodes {
		if connectedLookup[node.ID] {
			connected = append(connected, node)
		}
	}
	return connected
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

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	g := GraphDefinition{}
	if err := json.Unmarshal([]byte(request.Body), &g); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	graph, err := buildGraph(g)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	mGraph, err := json.Marshal(graph)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(mGraph),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
