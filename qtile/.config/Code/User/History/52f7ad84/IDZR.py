def isAnagram(s, t):
    if len(s) != len(t):
        return False
    contS, conutT = {}, {}
    for i in range(len(s)):
        contS[s[i]] = contS.get(s[i], 0) + 1
        conutT[t[i]] = conutT.get(t[i], 0) + 1
    return contS == conutT
