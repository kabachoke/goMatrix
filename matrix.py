from ctypes import sizeof
import numpy as np
from datetime import datetime

matr1 = np.random.rand(30000, 80)
matr2 = np.random.rand(80, 30000)
print(matr1)
start_time = datetime.now()
m = np.dot(matr1, matr2)
print(datetime.now() - start_time)

print(len(m), len(m[0]))
