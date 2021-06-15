import matplotlib.pyplot as plt

csa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
csa_time = [392.040, 299.721, 225.818, 165.427, 114.660, 73.818, 41.825, 18.803, 4.981, 1.431, 0.313, 0.269, 0.256]

fig, ax = plt.subplots()
fig.suptitle('Construction time', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Construction time, s', fontsize=12)

ax.plot(csa_len, csa_time, marker = 'o', c = 'r')
fig.savefig('construct_search.jpg')

plt.show()
