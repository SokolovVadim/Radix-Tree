import matplotlib.pyplot as plt

csa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
csa_mem = [2281, 2036, 1816, 1573, 1336, 1082, 817, 561, 309, 163, 45, 26, 17]

sa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
sa_mem = [5279, 4695, 4111, 3527, 2949, 2349, 1765, 1183, 599, 309, 60, 29, 9]

fig, ax = plt.subplots()
fig.suptitle('Lookup substring', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Heap allocation, kByte', fontsize=12)

ax.plot(csa_len, csa_mem, c = 'r', label = 'CSA')
ax.plot(sa_len, sa_mem, c = 'b', label = 'SA')

ax.legend(loc ="upper left")
fig.savefig('csa_sa.jpg')

plt.show()
