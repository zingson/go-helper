import rsa
# 生成密钥
(pubkey, privkey) = rsa.newkeys(1024)

# 保存密钥
#with open('public.pem','w+') as f:
#  f.write(pubkey.save_pkcs1().decode())
#with open('private.pem','w+') as f:
#  f.write(privkey.save_pkcs1().decode())

# 导入密钥
with open('public.pem','r') as f:
  pubkey = rsa.PublicKey.load_pkcs1(f.read().encode())
with open('private.pem','r') as f:
  privkey = rsa.PrivateKey.load_pkcs1(f.read().encode())

# *** 待签名内容
message = '123'

# ****签名
crypto_email_text = rsa.sign(message.encode(), privkey, 'SHA-1')
print("签名:")
print(crypto_email_text.hex())

# ****验签
vrs = rsa.verify(message.encode(), crypto_email_text, pubkey)
print(vrs)
