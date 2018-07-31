import influunt_core

from .executor import Executor

class Graph:
    def __init__(self):
        self._graph = influunt_core.new_graph()

    default_graph = None

    def __enter__(self):
        assert self.default_graph is None, "default graph is already initialized"
        Graph.default_graph = self
        return self

    def __exit__(self, *args):
        assert self.default_graph is self
        Graph.default_graph = None

    def add_op(self, spec):
        nodes = influunt_core.graph_add_op(self._graph, spec)
        wrappedNodes = []

        for node in nodes:
            wrappedNodes.append(Node(node))

        return wrappedNodes

    def executor(self):
        return Executor(self)


class Node:
    def __init__(self, node):
        self._node = node

    def __add__(self, anotherNode):
        return add(self, anotherNode)

    def __radd__(self, anotherNode):
        return add(anotherNode, self)

    def __getitem__(self, attr):
        return get_attr(self, attr)

    def __getattr__(self, attr):
        return get_attr(self, attr)

    def map(self, fn):
        return map(self, fn)

def map(list, fn, graph=None):
    if graph is None:
        graph = Graph.default_graph
    
    list = ensure_node(list)
    node = influunt_core.op_map(
        graph._graph, 
        list._node, 
        lambda item, i : fn(Node(item), Node(i))._node
    )
    return Node(node)
    

def add(a, b, graph=None):
    if graph is None:
        graph = Graph.default_graph

    a = ensure_node(a)
    b = ensure_node(b)

    return graph.add_op({
        "type": "Add",
        "inputs": [a._node, b._node]
    })[0]

def placeholder(graph=None):
    if graph is None:
        graph = Graph.default_graph
    
    return graph.add_op({
        "type": "Placeholder"
    })[0]

def const(val, graph=None):
    if graph is None:
        graph = Graph.default_graph

    return graph.add_op({
        "type": "Const",
        "attrs": {"value": val}
    })[0]


def cond(pred, a, b, graph=None):
    if graph is None:
        graph = Graph.default_graph
    
    if callable(a):
        a = a()

    if callable(b):
        b = b()

    pred = ensure_node(pred)
    a = ensure_node(a)
    b = ensure_node(b)

    return Node(influunt_core.op_cond(graph._graph, pred._node, a._node, b._node))

def parse_json(json, graph=None):
    if graph is None:
        graph = Graph.default_graph

    json = ensure_node(json)

    return graph.add_op({
        "type": "ParseJSON",
        "inputs": [json._node]
    })[0]

def get_attr(m, key, graph=None):
    if graph is None:
        graph = Graph.default_graph

    m = ensure_node(m)
    key = ensure_node(key)

    return graph.add_op({
        "type": "GetAttr",
        "inputs": [m._node, key._node]
    })[0]

def ensure_node(n, graph=None):
    if graph is None:
        graph = Graph.default_graph

    if isinstance(n, Node):
        return n
    else:
        return Node(influunt_core.op_const(graph._graph, n))
