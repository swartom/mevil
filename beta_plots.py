import numpy as np

from scipy.stats import beta

import matplotlib.pyplot as plt

fig, ax = plt.subplots(1, 1)


a, b = 2 , .75

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.99, a, b), 400)

ax.plot(x, beta.pdf(x, a, b),
        'r', lw=5,  label=f'{b:.2f}')

a, b = 2 , .25

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.8825, a, b), 100)

ax.plot(x, beta.pdf(x, a, b),
        '-',c="orange", lw=5, label=f'{b:.2f}')

a, b = 2 , .01

lb, ub = beta.support(a, b)

mean, var, skew, kurt = beta.stats(a, b, moments='mvsk')


x = np.linspace(beta.ppf(0.00, a, b),
                beta.ppf(0.9999, a, b), 1000)

ax.plot(x, beta.pdf(x, a, b),
        'g-', lw=5, label=f'{b:.2f}')

ax.set_xlim([0,1])
ax.set_ylim([0,6])

plt.xticks(size = 8)
plt.yticks(size = 8)

ax.grid(True,which="minor",ls=":",c=".1",alpha=.1)
ax.grid(True,which="major",ls="-",c=".1",alpha=.5)

ax.legend(loc='best', frameon=True,fontsize=10)

fig.tight_layout()
plt.savefig('data_pdf.pdf')
