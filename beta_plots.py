import numpy as np

from scipy.stats import beta

import matplotlib.pyplot as plt

fig, ax = plt.subplots(1, 1)
a, b = 2 , 1

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.99, a, b), 100)

ax.plot(x, beta.pdf(x, a, b),
        'r-', lw=5, label=f'{b:.2f}')

a, b = 2 , .75

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.99, a, b), 400)

ax.plot(x, beta.pdf(x, a, b),
        'c-', lw=5,  label=f'{b:.2f}')

a, b = 2 , .5

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.8825, a, b), 100)

ax.plot(x, beta.pdf(x, a, b),
        'b-', lw=5, label=f'{b:.2f}')

a, b = 2 , .25

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.6, a, b), 100,)

ax.plot(x, beta.pdf(x, a, b),
        'y-', lw=5, label=f'{b:.2f}')

a, b = 2 , .01

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.9999, a, b), 1000)

ax.plot(x, beta.pdf(x, a, b),
        'g-', lw=5, label=f'{b:.2f}')
ax.legend(loc='best', frameon=False,fontsize=20)
ax.set_xlim([0,1])
ax.set_ylim([0,6])
plt.xticks(size = 20)
plt.yticks(size = 20)

ax.grid(True,which="minor",ls="--",c="lightgray",alpha=.33)
ax.grid(True,which="major",ls=":",c="gray",alpha=1)
plt.show()
