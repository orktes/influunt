import influunt_core

class Executor:

    def __init__(self, graph):
        self.graph = graph
        self._executor = influunt_core.new_executor(graph._graph)

    def run(self, inputs, outputs):
        inputNodes = {}
        for key, value in inputs.items():
            inputNodes[key._node] = value

        outputNodes = []
        for node in outputs:
            outputNodes.append(node._node)

        return influunt_core.executor_run(self._executor, inputNodes, outputNodes)