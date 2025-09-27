#!/usr/bin/env python3

import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.ticker import (AutoMinorLocator, MultipleLocator)

G = nx.read_adjlist("test.adjlist")

g = G.to_undirected()



degree_sequence = sorted((d for n, d in g.degree()), reverse=True)
dmax = max(degree_sequence)



plt.rcParams["font.family"] = "monospace"



fig = plt.figure(2)

ax2 = fig.add_subplot(1,1,1)

degree_freq = nx.degree_histogram(g)
degrees = range(len(degree_freq))

ax2.loglog(degrees, degree_freq,c=".3",marker="x",ms=12,mew=3,lw =0.0)
# fig.suptitle(f'Degree Distribution of Lindenmayer Scale-free/Small-World Model (|V|={g.number_of_nodes():.1e})', fontsize = 35)
ax2.set_xlabel('Degree (Edge Count)',fontsize=30)
ax2.set_ylabel('Frequency (Number of Vertices)',fontsize=30)

plt.xticks(size = 35)
plt.yticks(size = 35)

ax2.grid(True,which="minor",ls="--",c="lightgray",alpha=.33)
ax2.grid(True,which="major",ls=":",c="gray",alpha=1)


plt.show()
