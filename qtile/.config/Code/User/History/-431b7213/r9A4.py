def isSubsequence(s: str, t: str) -> bool:
    if s or t == "":
        return False
    max_str = max(s, t)
    str_mat = {}

    for i in range(len(max_str)):
        if t[i] not in str_mat:
            str_mat[t[i]] = t[i]

    for i in range(len(s)):
        if s[i] not in str_mat:
            return False
    return True


s = "axc"
t = "ahbgdc"
print(isSubsequence(s, t))
