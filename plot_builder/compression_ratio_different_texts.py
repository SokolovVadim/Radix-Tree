import matplotlib.pyplot as plt

# average_ratio:  0.4052761480374943
# Amazon, DNA, Proteins, English, Pitches
csa_len =   [879, 782, 684, 586, 489, 391, 293, 196, 98, 49, 10, 5, 2]
csa_mem =   [1336, 1021, 1680, 1582, 1760] # 489 kB of text
sa_mem  =   [3707, 3627, 3627, 3627, 3628] # 489 kB of text
# radix_mem = [20248160, 20248160, 20248160, 20248160, 20248160]
# csa_mem_amazon = [2281, 2036, 1816, 1573, 1336, 1082, 817, 561, 309, 163, 45, 26, 17]
# csa_mem_dna = [19602, 17641, 15715, 16521, 14197, 11896, 9538, 7186, 4073, 2753, 1486, 493, 332, 246]

# sa_len = [879, 782, 684, 586, 498, 391, 293, 196, 98, 49, 10, 5, 2]
# sa_mem_amazon = [5279, 4695, 4111, 3527, 2949, 2349, 1765, 1183, 599, 309, 60, 29, 9]

ratio = []
average_ratio = 0
for i in range(len(csa_mem)):
    ratio.append(float(csa_mem[i] / sa_mem[i]))
    average_ratio += ratio[i]
    # print(ratio[i], " ")
average_ratio /= len(ratio)
print("average_ratio: ", average_ratio)
# plt.xticks(csa_mem, my_xticks)

fig, ax = plt.subplots()
fig.suptitle('Compression Ratio', fontsize=14)
plt.ylabel('Compression ratio', fontsize=12)

my_xticks = ['Amazon', 'DNA', 'Proteins', 'English', 'Pitches']
x_arr = [0, 1, 2, 3, 4]
plt.xticks(x_arr, my_xticks)
# plt.plot(csa_mem, ratio)
ax.scatter(x_arr, ratio, marker = 'o', c = 'r', label = 'CSA')

fig.savefig('compression_ratio_dif_texts.jpg')
plt.show()

