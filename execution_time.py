#!/usr/bin/env python3

x=[]

y=[]



with open("exec.txt") as f:
    for line in f:
        l = line.split(" ")
        x.append(int(l[0]))
        y.append(int(l[1]))

import matplotlib.pyplot as plt
plt.rcParams['text.usetex'] = True
plt.rcParams["font.family"] = "monospace"
fig = plt.figure(1)
ax2 = fig.add_subplot(1,1,1)
ax2.scatter(x,y,marker='.',color=".3",s=2)#,lw=0.0,ms=2)
ax2.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax2.grid(True,which="major",ls="-",c=".1",alpha=.5)
ax2.set_xlabel('Vertex Count',fontsize=10)
ax2.set_ylabel(r'Nanoseconds ($ns$)',fontsize=10)
plt.ticklabel_format(style='sci', axis='x', scilimits=(0,0))
plt.ticklabel_format(style='sci', axis='y', scilimits=(0,0))
plt.xticks(size = 8)
plt.yticks(size = 8)

fig.tight_layout()
plt.savefig('exec.pdf')
