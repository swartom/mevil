#!/usr/bin/env python3

import networkx as nx
import numpy as np
import matplotlib.pyplot as plt

G = nx.read_adjlist("test_2.adjlist")

g = G.to_undirected()
g.remove_edges_from(nx.selfloop_edges(g))

nodelist = list(g.nodes)

A = nx.to_numpy_array(g, nodelist=nodelist)


plt.rcParams['text.usetex'] = True
plt.rcParams["font.family"] = "monospace"
fig = plt.figure(1)
ax2 = fig.add_subplot(1,1,1)
ax2.imshow(A,cmap="Greys")


plt.ticklabel_format(style='sci', axis='x', scilimits=(0,0))
plt.ticklabel_format(style='sci', axis='y', scilimits=(0,0))
plt.xticks(size = 8)
plt.yticks(size = 8)

fig.tight_layout()
plt.savefig('heatmap.pdf')
