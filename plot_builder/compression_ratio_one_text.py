import matplotlib.pyplot as plt

csa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
csa_mem = [2281, 2036, 1816, 1573, 1336, 1082, 817, 561, 309, 163, 45, 26, 17]
csa_mem_dna = [19602, 17641, 15715, 16521, 14197, 11896, 9538, 7186, 4073, 2753, 1486, 493, 332, 246]

sa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
sa_mem = [5279, 4695, 4111, 3527, 2949, 2349, 1765, 1183, 599, 309, 60, 29, 9]

ratio = []
average_ratio = 0
for i in range(len(csa_len)):
    ratio.append(float(csa_mem[i] / sa_mem[i]))
    average_ratio += ratio[i]
    # print(ratio[i], " ")
average_ratio /= len(csa_len)
print("average_ratio: ", average_ratio)

fig, ax = plt.subplots()
fig.suptitle('Compression Ratio Amazon', fontsize=14)

plt.xlabel('Document size, Kbyte', fontsize=12)
plt.ylabel('Compression ratio', fontsize=12)

ax.scatter(csa_len, ratio, marker = 'o', c = 'r', label = 'CSA')
""" ax.plot(sa_len, sa_mem, marker = 'o', c = 'b', label = 'SA') """

""" ax.legend(loc ="upper left") """

fig.savefig('compression_ratio_amazon.jpg')

plt.show()

