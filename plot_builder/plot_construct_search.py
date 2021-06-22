import matplotlib.pyplot as plt
from scipy import interpolate
import numpy as np
# 3126 8009.034
csa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
csa_time = [392.040, 299.721, 225.818, 165.427, 114.660, 73.818, 41.825, 18.803, 4.981, 1.431, 0.313, 0.269, 0.256]
csa_len = sorted(csa_len)
csa_time = sorted(csa_time)

# interpolation
tck = interpolate.splrep(csa_len, csa_time)
xnew = np.arange(csa_len[0], csa_len[-1], 1.0)
ynew = interpolate.splev(xnew, tck, der=0)

fig, ax = plt.subplots()
fig.suptitle('Construction time Amazon', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Construction time, s', fontsize=12)

ax.scatter(csa_len, csa_time, marker = 'o', c = 'r')
ax.plot(xnew, ynew, c = 'b')
fig.savefig('construct_search.jpg')

plt.show()
