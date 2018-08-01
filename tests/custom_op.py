import sys
import random
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

randint = influunt.add_operation("random/randint", random.randint)

with influunt.Graph() as graph:
    start = influunt.placeholder()
    end = influunt.placeholder()

    randomNumber = randint(start, end)

    r = graph.executor().run({start: 1, end: 100}, [randomNumber])
 
    print(r)
    assert r[0] >= 1 and r[0] <= 100
