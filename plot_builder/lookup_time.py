import matplotlib.pyplot as plt
import numpy as np

csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98]
csa_time = [157.230, 144.460, 110.304, 137.585, 97.767, 87.997, 86.016, 83.449, 61.127, 50.558]
sa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98]
sa_time = [12.251, 11.857, 12.331, 12.969, 12.777, 12.680, 11.795, 11.555, 11.752, 11.657]
radix_len = [20, 50, 60, 70, 80, 90, 100, 120, 150, 200, 300, 400, 500, 600, 700, 800, 900]
radix_time = [0.220, 0.242, 0.243, 0.237, 0.243, 0.244, 0.243, 0.245, 0.256, 0.258, 0.261, 0.263, 0.267, 0.268, 0.270, 0.274, 0.277]

fig, ax = plt.subplots()
fig.suptitle('Lookup', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Lookup time, s', fontsize=12)

ax.scatter(csa_len, csa_time, marker = 'o', c = 'r', label = 'CSA')
ax.scatter(sa_len, sa_time, marker = 'o', c = 'b', label = 'SA')
ax.scatter(radix_len, radix_time, marker = 'o', c = 'g', label = 'Radix')
ax.legend(bbox_to_anchor=(1.05, 0.6),  fancybox=True, shadow=True)
fig.savefig('Lookup_time.jpg')

plt.show()
