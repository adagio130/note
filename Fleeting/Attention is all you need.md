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
2. 