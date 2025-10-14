#!/usr/bin/env python3

x=[]

y=[]



with open("exec.txt") as f:
    for line in f:
        l = line.split(" ")
        edges = int(l[0]) * 60
        x.append(int(l[0]))
        # y.append(int(l[1])/10**6)
        y.append((edges/6)/(float(l[1])/10**9)/10**6)

print(y)

import matplotlib.pyplot as plt
plt.rcParams['text.usetex'] = True
plt.rcParams["font.family"] = "monospace"
fig = plt.figure(1)
ax2 = fig.add_subplot(1,1,1)
ax2.semilogy(x,y,marker='x',color="red",lw=1,ms=2)
ax2.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax2.grid(True,which="major",ls="-",c=".1",alpha=.5)
ax2.set_xlabel('Vertex Count',fontsize=10)
ax2.set_ylabel('Million Edges Per Processor Per Second ($mE/PE/s$)',fontsize=10)
ax2.set_ylim(-100,10**4)

plt.xticks(size = 8)
plt.yticks(size = 8)

fig.tight_layout()
plt.savefig('exec.pdf')
