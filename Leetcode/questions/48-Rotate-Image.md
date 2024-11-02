1. forloop len(matrix) - 2 until 0
	1. With each iteration, we move one layer inward. we move totally **math.Floor((n+1) / 2)** times
	2. **n + 1** because when the length of matrix is odd, the core point is only one cell, so we need to plus one
	3. **module 2** because each iteration would narrow down by two lines
2. in each layer, the four cell need to replace literately
	1.  