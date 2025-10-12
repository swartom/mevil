#!/usr/bin/env python3
import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.ticker import (AutoMinorLocator, MultipleLocator)
l = list()
for f in open("exec.txt"):
    l.append(int(f.split(" ")[1]))

G = nx.read_adjlist("test_2.adjlist")

g = G.to_undirected()
print(f'{((len(g.edges)/10)/(l[0]/10**9))/10**6:.2f}m/PE/s')

degree_sequence = sorted((d for n, d in g.degree()), reverse=True)
dmax = max(degree_sequence)



plt.rcParams["font.family"] = "monospace"

fig = plt.figure(2)

ax2 = fig.add_subplot(1,1,1)
# axins = ax2.inset_axes(
#     [0.55, 0.15, 0.4, 0.2],
#     xlim=(10^1 , 20), ylim=(10**3 * 2, 10**4  * 3))

degree_freq = nx.degree_histogram(g)
degrees = range(len(degree_freq))

ax2.loglog(degrees, degree_freq,c="#B4111B",marker=".",ms=2.5,mew=.2,lw =0.0, label=f'p=0.75')
# axins.loglog(degrees, degree_freq,c="#B4111B",marker=".",ms=2.5,mew=.2,lw =.5, label=f'p=0.75')

# G = nx.read_adjlist("test.adjlist")

# g = G.to_undirected()

# degree_sequence = sorted((d for n, d in g.degree()), reverse=True)
# dmax = max(degree_sequence)

# degree_freq = nx.degree_histogram(g)
# degrees = range(len(degree_freq))

# ax2.loglog(degrees, degree_freq,c="orange",marker="x",ms=2,mew=1,lw =0.0, label=f'p=0.25')
# # axins.loglog(degrees, degree_freq,c="orange",marker="x",ms=2,mew=1,lw =.5, label=f'p=0.25')

# G = nx.read_adjlist("test_1.adjlist")

# g = G.to_undirected()



# degree_sequence = sorted((d for n, d in g.degree()), reverse=True)
# dmax = max(degree_sequence)

# degree_freq = nx.degree_histogram(g)
# degrees = range(len(degree_freq))

# ax2.loglog(degrees, degree_freq,c="green",marker="^",ms=1.5,mew=.2,lw =0.0, label=f'p=0.01')
# # axins.loglog(degrees, degree_freq,c="green",marker="^",ms=1.5,mew=.2,lw =.5, label=f'p=0.01')

# fig.suptitle(f'Degree Distribution of Lindenmayer Scale-free/Small-World Model (|V|={g.number_of_nodes():.1e})', fontsize = 35)
ax2.set_xlabel('Degree (Edge Count)',fontsize=10)
ax2.set_ylabel('Frequency (Number of Vertices)',fontsize=10)

plt.xticks(size = 8)
plt.yticks(size = 8)

# axins.get_xaxis().set_visible(False)
# axins.get_yaxis().set_visible(False)

ax2.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax2.grid(True,which="major",ls="-",c=".1",alpha=.5)


# ax2.indicate_inset_zoom(axins, edgecolor="black")

# ax2.legend(loc='best', frameon=True,fontsize=10)

fig.tight_layout()
plt.savefig('data_pdf.pdf')
