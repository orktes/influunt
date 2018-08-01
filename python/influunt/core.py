import influunt_core

from .executor import Executor

class Graph:
    def __init__(self, graph=None):
        if graph is None:
            graph = influunt_core.new_graph()

        self._graph = graph

    default_graph = None

    def __enter__(self):
        assert self.default_graph is None, "default graph is already initialized"
        Graph.default_graph = self
        return self

    def __exit__(self, *args):
        assert self.default_graph is self
        Graph.default_graph = None

    def node_by_name(self, name, index=0):
        node = influunt_core.graph_node_by_name(self._graph, name, index)
        if node is None:
            return None

        return Node(node)

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
        self._name = influunt_core.node_get_name(node)

    def __repr__(self):
        return 'Node(name=%s)' % (self._name)

    def __add__(self, anotherNode):
        return add(self, anotherNode)

    def __radd__(self, anotherNode):
        return add(anotherNode, self)

    def __sub__(self, anotherNode):
        return sub(self, anotherNode)

    def __rsub__(self, anotherNode):
        return sub(anotherNode, self)

    def __mod__(self, anotherNode):
        return mod(self, anotherNode)

    def __rmod__(self, anotherNode):
        return mod(anotherNode, self)

    def __truediv__(self, anotherNode):
        return sub(self, anotherNode)

    def __rtruediv__(self, anotherNode):
        return sub(anotherNode, self)

    def __lt__(self, anotherNode):
        return less_than(self, anotherNode)

    def __le__(self, anotherNode):
        return less_or_equal(self, anotherNode)

    def __gt__(self, anotherNode):
        return greater_than(self, anotherNode)

    def __ge__(self, anotherNode):
        return greater_or_equal(self, anotherNode)

    def __eq__(self, anotherNode):
        return equal(self, anotherNode)

    def __req__(self, anotherNode):
        return equal(anotherNode, self)

    def __getitem__(self, attr):
        return get_attr(self, attr)

    def __getattr__(self, attr):
        return get_attr(self, attr)

    def __hash__(self):
        return hash(self._name)
        
    def get_attr(self, key):
        return get_attr(self, attr)

    def map(self, fn):
        return map(self, fn)

def load_graph(filepath):
    return Graph(influunt_core.read_graph_from_file(filepath))

def save_graph(graph, filepath):
    return influunt_core.write_graph_to_file(graph._graph, filepath)

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
    
def generate_binary_op(type):
    def op(a, b, graph=None):
        if graph is None:
            graph = Graph.default_graph

        a = ensure_node(a)
        b = ensure_node(b)

        return graph.add_op({
            "type": type,
            "inputs": [a._node, b._node]
        })[0]

    return op


add = generate_binary_op("Add")
sub = generate_binary_op("Sub")
mod =  generate_binary_op("Mod")
mul =  generate_binary_op("Mul")
mod =  generate_binary_op("Div")
less_than = generate_binary_op("LessThan")
less_or_equal = generate_binary_op("LessOrEqual")
greater_than = generate_binary_op("GreaterThan")
greater_or_equal = generate_binary_op("GreaterOrEqual")
equal = generate_binary_op("Equal")
not_equal = generate_binary_op("NotEqual")

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

    return graph.add_op({
        "type": "Cond",
        "inputs": [pred._node, a._node, b._node]
    })[0]

def parse_json(json, graph=None):
    if graph is None:
        graph = Graph.default_graph

    json = ensure_node(json)

    return graph.add_op({
        "type": "ParseJSON",
        "inputs": [json._node]
    })[0]

get_attr = generate_binary_op("GetAttr")

def ensure_node(n, graph=None):
    if graph is None:
        graph = Graph.default_graph

    if isinstance(n, Node):
        return n
    else:
        return const(n, graph)
