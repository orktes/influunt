import sys
import random
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

randint = influunt.add_operation("randint", random.randint)

with influunt.Graph() as graph:
    randomNumber = randint(1, 100)

    r = graph.executor().run({}, [randomNumber])
 
    print(r)
    assert r[0] >= 1 and r[0] <= 100
