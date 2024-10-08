#machine-learning

![[Pasted image 20240326135310.png]]
1. Self Attention
2. Multi-Head Attention
3. Position Encoding
4. Encoder
5. Decoder
6. Feed-Forward Networks
7. Layer Normalization and Residual Connection

steps:

prepare data:
1. tokenize (斷詞)
2. using pre-train model to convert each character to word embedding
3. add special annotation such as \[CLS\]、\[SEP\]( for start、end)
4. add positional encoding

transformer
1. input word embedding and positional vector to generate query, key, value for each word.
	1. Q =  $W^Q \cdot I$
	2. K =  $W^K \cdot I$
	3. V =  $W^V \cdot I$
2. calculate the attention score
	1. dot the Q and K
		1. show the Q how important is for each K
		2. divide $\sqrt{d}$<sub>k</sub>
3. 把attention score跟value做加權求和
4. above is called head attention
5. it would do multiple times parallelize which called multi-head attention
6. encoder output key and value as decoder's key and value