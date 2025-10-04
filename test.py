#!/usr/bin/env python3

import networkx as nx
import matplotlib.pyplot as plt
G = nx.read_adjlist("test_2.adjlist")
print("readfile")
g = G.to_undirected()
print("undirected")

# matrix = nx.convert_matrix(g)

nx.draw(g,node_color="#B4111B",node_size=2,edge_color=".2")
# plt.imshow(matrix, cmap='hot', interpolation='nearest')
# plt.show()

plt.savefig('graph.pdf')
