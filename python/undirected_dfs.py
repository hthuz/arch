
def undirectedIsCyclic(graph, n):
	visited = [0] * n
	rec_stack = [0] * n
	for node in range(0, n):
		if visited[node] == 0:
			if _undrectedIsCyclic(graph, node, node, visited, rec_stack) == True:
				return True
	return False

def _undrectedIsCyclic(graph, node, parent, visited, rec_stack):
	print(node, parent)
	visited[node] = 1
	rec_stack[node] = 1
	for adj_node in graph[node]:
		if visited[adj_node] == 0:
			if _undrectedIsCyclic(graph, adj_node, node, visited, rec_stack) == True:
				return True
		if adj_node != parent and rec_stack[adj_node] == 1:
			return True
	rec_stack[node] = 0

	return False

def edge_to_graph(edges: list, n: int):
	graph = [[] for _ in range(n)]
	for (node1, node2) in edges:
		graph[node1].append(node2)
		graph[node2].append(node1)
	return graph


if __name__ == "__main__":
	edges = [(0,1),(0,2),(1,3),(2,3)]
	graph = edge_to_graph(edges, 4)
	print(undirectedIsCyclic(graph, 4))


