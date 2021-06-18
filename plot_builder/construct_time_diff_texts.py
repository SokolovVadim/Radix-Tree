import matplotlib.pyplot as plt
import numpy as np

# Amazon, DNA, Proteins, English, Pitches

# csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 49, 10, 5, 2] # dna
# csa_time = [486.965, 385.449, 301.335, 230.011, 167.257, 115.855, 73.921, 41.823, 18.856, 4.920, 0.302] # english
# csa_time = [491.614, 385.973, 309.387, 232.106, 168.735, 116.836, 73.656, 42.861, 18.788, 4.922, 1.446, 0.304, 0.267, 0.255] # dna
# csa_time = [480.690, 386.661, 311.412, 229.365, 166.962, 116.043, 74.501, 42.195, 19.111, 5.021, 0.312] # proteins
csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 10]
csa_time = [392, 386, 387, 385, 382] # 879 kB
sa_time = [0.333, 0.346, 0.345, 0.342, 0.348]
radix_time = [1335, 1306, 1315, 1331, 1326]
#csa_time = [478.944, 382.251, 300.100, 228.356, 172.207, 116.840, 75.856, 42.822, 19.191, 4.936, 0.303] # pitches
# csa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2] # amazon
# csa_time = [392.040, 299.721, 225.818, 165.427, 114.660, 73.818, 41.825, 18.803, 4.981, 1.431, 0.313, 0.269, 0.256] # amazon

fig, ax = plt.subplots()
fig.suptitle('Construction time', fontsize=14)

#plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Construction time, s', fontsize=12)

my_xticks = ['Amazon', 'DNA', 'Proteins', 'English', 'Pitches']
x_arr = [0, 1, 2, 3, 4]
plt.xticks(x_arr, my_xticks)

ax.scatter(x_arr, csa_time, marker = 'o', c = 'r', label = 'CSA')
ax.scatter(x_arr, sa_time, marker = 'o', c = 'b', label = 'SA')
ax.scatter(x_arr, radix_time, marker = 'o', c = 'g', label = 'Radix')
ax.legend(bbox_to_anchor=(1.1, 1.05),  fancybox=True, shadow=True)
fig.savefig('construct_time_diff_texts.jpg')

plt.show()
