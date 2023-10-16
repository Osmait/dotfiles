def isSubsequence(s: str, t: str) -> bool:
    max_str = max(s, t)
    str_mat = {}

    for i in max_str:
        if i not in str_mat:
            str_mat[i] = i
        if s[i] not in str_mat:
            return False
    return True


s = "axc"
t = "ahbgdc"
print(isSubsequence(s, t))
