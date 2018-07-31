import sys
import os
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

with influunt.Graph() as graph:
    p = influunt.placeholder()
    c = p + 1 + 123 + 231

    res = graph.executor().run({p:1}, [c])
    assert res == [356]

    influunt.save_graph(graph, "test.graph")

with influunt.load_graph("test.graph") as graph:
    input = graph.node_by_name("Placeholder:0")
    output = graph.node_by_name("Add:6")

    res = graph.executor().run({input:1}, [output])
    assert res == [356]

os.remove("test.graph")