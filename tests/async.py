import sys
import random
import time
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

with influunt.Graph() as graph:
    a = influunt.placeholder()
    b = influunt.placeholder()

    sum = a + b

    def callback(sum):
        print(sum)
        assert sum == 101

    graph.executor().run_async({a: 1, b: 100}, [sum], callback)
 
time.sleep(0.5)