def isSubsequence(s: str, t: str) -> bool:
    max_str = max(s, t)

    for i in max_str:
        print(i)


s = "abc"
t = "ahbgdc"
print(isSubsequence(s, t))
