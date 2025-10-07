#!/usr/bin/env python3



import networkx as nx
import numpy as np
import matplotlib.pyplot as plt

G = nx.read_adjlist("test_2.adjlist")
g = G

g.remove_edges_from(nx.selfloop_edges(g))

plt.rcParams["font.family"] = "monospace"

A = nx.to_numpy_array(g)

fig = plt.figure(1)

plt.imshow(A,cmap="Greys",interpolation='none')
plt.xticks(size = 8)
plt.yticks(size = 8)
plt.grid(True,which="minor",ls=":",c=".1",alpha=.1)
plt.grid(True,which="major",ls="-",c=".1",alpha=.5)

fig.tight_layout()

plt.savefig('heatmap.pdf')
