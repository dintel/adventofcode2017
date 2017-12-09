package main

import (
	"log"
	"strconv"
	"strings"
)

type Node struct {
	Name        string
	Weight      int
	TotalWeight int
	Children    []*Node
	Parent      *Node
}

func fillTotalWieghts(node *Node) {
	node.TotalWeight = node.Weight
	if len(node.Children) == 0 {
		return
	}
	for _, child := range node.Children {
		fillTotalWieghts(child)
		node.TotalWeight += child.TotalWeight
	}
}

func findUnbalancedNode(node *Node) *Node {
	weights := make(map[int][]*Node)
	for _, child := range node.Children {
		weights[child.TotalWeight] = append(weights[child.TotalWeight], child)
	}
	if len(weights) == 1 {
		return nil
	}
	for _, nodes := range weights {
		if len(nodes) == 1 {
			result := findUnbalancedNode(nodes[0])
			if result == nil {
				return node
			}
			return result
		}
	}
	return nil
}

func Day7(part int, data []byte) {
	lines := strings.Split(string(data), "\n")
	nodes := make(map[string]*Node)
	for i, line := range lines {
		fields := strings.Fields(line)
		name := fields[0]
		weight, err := strconv.Atoi(strings.Trim(fields[1], "()"))
		if err != nil {
			log.Fatalf("Failed parsing weight at line %d", i)
		}
		if nodes[name] != nil {
			nodes[name].Weight = weight
		} else {
			nodes[name] = &Node{Name: name, Weight: weight}
		}
		if len(fields) > 2 {
			for _, child := range fields[3:] {
				child = strings.Trim(child, ",")
				if nodes[child] == nil {
					nodes[child] = &Node{Name: child}
				}
				nodes[name].Children = append(nodes[name].Children, nodes[child])
				nodes[child].Parent = nodes[name]
			}
		}
	}

	log.Printf("Loaded tree with %d nodes", len(nodes))
	var root *Node
	for _, node := range nodes {
		if node.Parent == nil {
			root = node
		}
	}
	switch part {
	case 1:
		log.Printf("root node name is %s", root.Name)
	case 2:
		fillTotalWieghts(root)
		unbalancedNode := findUnbalancedNode(root)
		weights := make(map[int][]*Node)
		for _, child := range unbalancedNode.Children {
			weights[child.TotalWeight] = append(weights[child.TotalWeight], child)
		}
		balancedChildWeight := 0
		unbalancedChild := (*Node)(nil)
		for weight, nodes := range weights {
			if len(nodes) != 1 {
				balancedChildWeight = weight
			} else {
				unbalancedChild = nodes[0]
			}
		}
		delta := unbalancedChild.TotalWeight - balancedChildWeight
		log.Printf("unbalanced node weight is %d, but should be %d", unbalancedChild.Weight, unbalancedChild.Weight-delta)
	}
}
