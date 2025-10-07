#!/usr/bin/env python3

import networkx as nx
import numpy as np
import matplotlib.pyplot as plt

l = [(10**4)//2,(10**4)//2]

import time


now = time.time()

nx.stochastic_block_model(l,[[0.8,0.5],[0.5,0.8]])

then = time.time()

print(then-now)
