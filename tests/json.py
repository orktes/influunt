import sys
from os.path import dirname, join

sys.path.append(join(dirname(__file__), "../build/sharedlib/linux/amd64"))
sys.path.append(join(dirname(__file__), "../python"))

import influunt

with influunt.Graph() as graph:
    m = influunt.parse_json('{"foo": {"bar": "biz"}, "biz": ["foo", "bar"]}')
    biz = m.foo.bar
    fooboo = m[biz].map(lambda item, i : item + biz)

    r = graph.executor().run({}, [biz, fooboo])
    print(r)
    assert r == ("biz", ['foobiz', 'barbiz'])