import matplotlib.pyplot as plt
import numpy as np
from scipy import interpolate

# csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 49, 10, 5, 2] # dna
# csa_time = [486.965, 385.449, 301.335, 230.011, 167.257, 115.855, 73.921, 41.823, 18.856, 4.920, 0.302] # english
# csa_time = [491.614, 385.973, 309.387, 232.106, 168.735, 116.836, 73.656, 42.861, 18.788, 4.922, 1.446, 0.304, 0.267, 0.255] # dna
csa_time = [480.690, 386.661, 311.412, 229.365, 166.962, 116.043, 74.501, 42.195, 19.111, 5.021, 0.312] # proteins
csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 10]
#csa_time = [478.944, 382.251, 300.100, 228.356, 172.207, 116.840, 75.856, 42.822, 19.191, 4.936, 0.303] # pitches

csa_len = sorted(csa_len)
csa_time = sorted(csa_time)

# interpolation
tck = interpolate.splrep(csa_len, csa_time)
xnew = np.arange(csa_len[0], csa_len[-1] + 70, 80.0)
ynew = interpolate.splev(xnew, tck, der=0)

array = np.polyfit(csa_len, np.log(csa_time), 1)
# print(array)
interp = []
for i in range(len(csa_len)):
    interp.append(np.exp(array[1]) * np.exp(array[0] * csa_len[i]))
fig, ax = plt.subplots()
fig.suptitle('Construction time Proteins', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Construction time, s', fontsize=12)

ax.scatter(csa_len, csa_time, marker = 'o', c = 'r')
ax.plot(xnew, ynew, c = 'b')
fig.savefig('construct_time_Proteins.jpg')

plt.show()
