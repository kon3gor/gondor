package gondor

import (
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"golang.org/x/xerrors"
)

func Parse(into interface{}, base string, layers ...string) error {
	rootNode, err := parse(base)
	if err != nil {
		return nil
	}

	for _, layer := range layers {
		node, err := parse(layer)
		if err != nil {
			return err
		}

		applyLayer(rootNode, node)
	}

	result, err := rootNode.MarshalYAML()
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, into)
}

func parse(configFile string) (ast.Node, error) {
	// NOTE(kon3gor): Use mode = 0 to ignore comments
	f, err := parser.ParseFile(configFile, 0)
	if err != nil {
		return nil, err
	}

	// NOTE(kon3gor): For now we assume that there is only one doc
	if len(f.Docs) != 1 {
		return nil, xerrors.Errorf("wrong number of documents in file: %d", len(f.Docs))
	}

	return f.Docs[0].Body, nil
}

func applyLayer(node, layer ast.Node) error {
	if node.Type() != layer.Type() {
		return xerrors.Errorf("node types doesn't match. %s vs %s", node.Type().String(), layer.Type().String())
	}

	visit(node, layer)

	return nil
}

func visit(node, layer ast.Node) {
	switch node.Type() {
	case ast.MappingType:
		visitMapping(node.(*ast.MappingNode), layer.(*ast.MappingNode))
	case ast.MappingValueType:
		visitMappingValue(node.(*ast.MappingValueNode), layer.(*ast.MappingValueNode))
	case ast.SequenceType:
		visitArray(node.(*ast.SequenceNode), layer.(*ast.SequenceNode))
	case ast.StringType:
		visitString(node.(*ast.StringNode), layer.(*ast.StringNode))
	case ast.IntegerType:
		visitInteger(node.(*ast.IntegerNode), layer.(*ast.IntegerNode))
	case ast.BoolType:
		visitBoolean(node.(*ast.BoolNode), layer.(*ast.BoolNode))
	case ast.FloatType:
		visitFloat(node.(*ast.FloatNode), layer.(*ast.FloatNode))
	}
}

func visitMapping(node, layer *ast.MappingNode) {
	valueToKey := make(map[string]ast.Node)
	for _, value := range node.Values {
		valueToKey[value.Key.String()] = value
	}

	for _, value := range layer.Values {
		nodeValue, ok := valueToKey[value.Key.String()]
		if !ok {
			node.Values = append(node.Values, value)
			continue
		}

		applyLayer(nodeValue, value)
	}
}

func visitMappingValue(node, layer *ast.MappingValueNode) {
	applyLayer(node.Value, layer.Value)
}

func visitArray(node, layer *ast.SequenceNode) {
	// TODO(kon3gor): Think of a smarter way to merge arrays
	node.Values = layer.Values
}

func visitString(node, layer *ast.StringNode) {
	node.Value = layer.Value
}

func visitInteger(node, layer *ast.IntegerNode) {
	node.Value = layer.Value
}

func visitBoolean(node, layer *ast.BoolNode) {
	node.Value = layer.Value
}

func visitFloat(node, layer *ast.FloatNode) {
	node.Value = layer.Value
}
