import matplotlib.pyplot as plt
import numpy as np

# Amazon, DNA, Proteins, English, Pitches
csa_len = [977, 879, 782, 684, 586, 489, 391, 293, 196, 98, 10]
csa_time = [50.558, 27.652, 38.759, 70.454, 47.331] # 100 kB text, 80 kB substring
sa_time = [11.657, 12.035, 11.987, 11.812, 11.936]
radix_time = [0.243, 0.242, 0.242, 0.242, 0.243]

fig, ax = plt.subplots()
fig.suptitle('Lookup', fontsize=14)

plt.ylabel('Lookup time, s', fontsize=12)

my_xticks = ['Amazon', 'DNA', 'Proteins', 'English', 'Pitches']
x_arr = [0, 1, 2, 3, 4]
plt.xticks(x_arr, my_xticks)

ax.scatter(x_arr, csa_time, marker = 'o', c = 'r', label = 'CSA')
ax.scatter(x_arr, sa_time, marker = 'o', c = 'b', label = 'SA')
ax.scatter(x_arr, radix_time, marker = 'o', c = 'g', label = 'Radix')
ax.legend(bbox_to_anchor=(1.1, 1.05),  fancybox=True, shadow=True)
fig.savefig('lookup_diff_texts.jpg')

plt.show()