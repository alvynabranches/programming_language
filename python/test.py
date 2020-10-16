import pandas as pd
import numpy as np

c = np.linspace(0, 224, 224)
c = np.round(c)
c = np.array(c, dtype=int)
print(c)
df = pd.read_csv('libraries.csv', names=c)
print(df.T.sort_values(by=0))
# df.T.sort_values(by=0).to_csv('python_libraries.csv', index=False)