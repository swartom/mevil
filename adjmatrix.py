#!/usr/bin/env python3

import networkx as nx
import numpy as np
import matplotlib.pyplot as plt

G = nx.read_adjlist("test_2.adjlist")
g = G

g.remove_edges_from(nx.selfloop_edges(g))



A = nx.to_numpy_array(g)

fig = plt.figure(1)

plt.imshow(A,cmap="Greys",interpolation='none')



plt.savefig('heatmap.pdf')
