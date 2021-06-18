import matplotlib.pyplot as plt
import numpy as np
from scipy import interpolate

# csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 49, 10, 5, 2] # dna
# csa_time = [486.965, 385.449, 301.335, 230.011, 167.257, 115.855, 73.921, 41.823, 18.856, 4.920, 0.302] # english
# csa_time = [491.614, 385.973, 309.387, 232.106, 168.735, 116.836, 73.656, 42.861, 18.788, 4.922, 1.446, 0.304, 0.267, 0.255] # dna
csa_time = [157.230, 144.460, 110.304, 137.585, 97.767, 87.997, 86.016, 83.449, 61.127, 50.558]
csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98]
#csa_time = [478.944, 382.251, 300.100, 228.356, 172.207, 116.840, 75.856, 42.822, 19.191, 4.936, 0.303] # pitches

fig, ax = plt.subplots()
fig.suptitle('Lookup', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Lookup time, s', fontsize=12)

ax.scatter(csa_len, csa_time, marker = 'o', c = 'r')
fig.savefig('Lookup_time_csa_amazon.jpg')

plt.show()
