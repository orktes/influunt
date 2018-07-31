import sys
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

with influunt.Graph() as graph:
    p = influunt.placeholder()
    n = p + 2
    n = 2 + n
    n += 1
    l = influunt.const([1, 2, 3])
    mapped = l.map(lambda x, i : x + 1 + i)
    foo = influunt.cond(True, "foo", "bar")
    bar = influunt.cond(False, lambda : "foo", lambda : "bar")

    r = graph.executor().run({p: 1}, [n, mapped, foo, bar])
 
    print(r)
    assert r == [6, [2, 4, 6], "foo", "bar"]
    

