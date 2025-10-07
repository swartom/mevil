#!/usr/bin/env python3
import networkx as nx
import numpy as np
import matplotlib.pyplot as plt




x=[]

y=[]

count = 5
samples = 40

points = [ (int((10**(i+1))/samples))*(j) for i in range(count) for j in range(samples)]
points = [i for i in points if i > 10]


print(points)
import time

repeats = 2
for point in points:
    print(point)
    now = time.time_ns()
    for i in range(repeats):
        nx.barabasi_albert_graph(point,1)

    then = time.time_ns()

    x.append(point)
    y.append((then-now)/repeats)

import matplotlib.pyplot as plt
plt.rcParams['text.usetex'] = True
plt.rcParams["font.family"] = "monospace"
fig = plt.figure(1)
ax2 = fig.add_subplot(1,1,1)
ax2.plot(x,y,marker='.',color="red")

x=[]

y=[]



with open("exec.txt") as f:
    for line in f:
        l = line.split(" ")
        x.append(int(l[0]))
        y.append(int(l[1]))

ax2.plot(x,y,marker='.',color=".3")#,lw=0,ms=2)
ax2.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax2.grid(True,which="major",ls="-",c=".1",alpha=.5)
ax2.set_xlabel('Vertex Count',fontsize=10)
ax2.set_ylabel(r'Nanoseconds ($ns$)',fontsize=10)
# plt.ticklabel_format(style='sci', axis='x', scilimits=(0,0))
# plt.ticklabel_format(style='sci', axis='y', scilimits=(0,0))
plt.xticks(size = 8)
plt.yticks(size = 8)

fig.tight_layout()
plt.savefig('exec.pdf')
