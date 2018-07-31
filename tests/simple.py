import sys
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
import influunt_core

graph = influunt_core.new_graph()
somerandomlist = influunt_core.op_const(graph, [1,2,3])

trueVal = influunt_core.op_const(graph, True)
falseVal = influunt_core.op_const(graph, False)

a = influunt_core.op_placeholder(graph)
b = influunt_core.op_const(graph, 321)
c = influunt_core.op_add(graph, a, b)

[manualOpRes] = influunt_core.graph_add_op(graph, {
    "type": "Add",
    "inputs": [a, c]
})

randomListMap = influunt_core.op_map(graph, somerandomlist, lambda x, i : influunt_core.op_add(graph, x, i))

json = influunt_core.op_const(graph, '{"foo": "bar", "biz": "foo"}')
m = influunt_core.op_parsejson(graph, json)
foo = influunt_core.op_getattr(graph, m, influunt_core.op_const(graph, "foo"))
biz = influunt_core.op_getattr(graph, m, influunt_core.op_const(graph, "biz"))

isFoo = influunt_core.op_cond(graph, trueVal, foo, biz)
notFoo = influunt_core.op_cond(graph, falseVal, foo, biz)

executor = influunt_core.new_executor(graph)
res = influunt_core.executor_run(executor, {a: 123}, [b, c, manualOpRes, randomListMap, trueVal, falseVal, json, m, isFoo, notFoo])

assert res == [321, 444, 567, [1, 3, 5], False, True, '{"foo": "bar", "biz": "foo"}', {'biz': 'foo', 'foo': 'bar'}, 'foo', 'bar']
print(res)