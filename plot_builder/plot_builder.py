import matplotlib.pyplot as plt

radix_len = [120, 110, 100, 90, 80, 70, 60, 50, 40, 30, 20]
radix_time = [82.941, 71.651, 60.559, 47.864, 37.66, 29.485, 21.827, 14.785, 9.637, 5.537, 2.694]

suffix_len = [19, 24, 37, 52, 66, 81, 90, 120]
suffix_time = [0.241, 0.242, 0.244, 0.245, 0.246, 0.248, 0.250, 0.251]

fig, ax = plt.subplots()
fig.suptitle('Lookup substring', fontsize=14)

plt.xlabel('Document size, Kbytes', fontsize=12)
plt.ylabel('Time, s', fontsize=12)

ax.plot(radix_len, radix_time, c = 'r', label = 'Radix tree')
ax.plot(suffix_len, suffix_time, c = 'b', label = 'Suffix array')

ax.legend(loc ="upper left")
fig.savefig('radix.jpg')

plt.show()
