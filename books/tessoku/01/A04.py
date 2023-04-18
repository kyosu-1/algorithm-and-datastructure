N = int(input())

# Nを2進法に変換した値を10桁で出力する
print(format(N, 'b').zfill(10))
# format(N, 'b')はNを2進法に変換した値を文字列で返す
# zfill(10)は文字列の左側を0で埋めて10桁にする

# ちゃんと計算をして答える
# 9から0でループを回す
# ans = ""
# for i in range(9, -1, -1):
#     wari = 1 << i
#     ans += str((N // wari) % 2)
# 
# print(ans)