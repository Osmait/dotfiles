def isSubsequence(s: str, t: str) -> bool:
    max_str = max(s, t)
    str_mat = {}

    for i in max_str:
        if i not in str_mat:
            str_mat[i] = i


s = "abc"
t = "ahbgdc"
print(isSubsequence(s, t))
