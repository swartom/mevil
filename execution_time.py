#!/usr/bin/env python3

x=[]

y=[]



with open("exec.txt") as f:
    for line in f:
        l = line.split(" ")
        x.append(int(l[0]))
        y.append(int(l[1]))

import matplotlib.pyplot as plt

plt.rcParams["font.family"] = "monospace"

fig = plt.figure(1)
ax2 = fig.add_subplot(1,1,1)
ax2.scatter(x,y,marker='.',color=".3",s=1)
ax2.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax2.grid(True,which="major",ls="-",c=".1",alpha=.5)


fig.tight_layout()
plt.savefig('exec.pdf')
